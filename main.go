package main


import (
	//"context"
	"net/http"
	"crypto/subtle"
	"os"
	"github.com/go-kit/kit/auth/jwt"
	stdjwt "github.com/dgrijalva/jwt-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
			//customersvc "salv_prj/customerservice"
			//ordersvc "salv_prj/orderservice"
			//partnersvc "salv_prj/partnerservice"
	usersvc "salv_prj/user"
	schoolsvc "salv_prj/school"
	"salv_prj/auth"
	"salv_prj/store"
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"

)

var loggrer = logrus.New()

func main() {

	//store.ConfigureApp("staging")
	defer store.Database.Close()

	//ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
	loggrer.Out = os.Stdout

	//loggrer.Formatter = &logrus.JSONFormatter{}
	//logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//logrus.SetLevel(logrus.WarnLevel)
	//kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("SigningString"), nil }

//jwt

	key := []byte("supersecret")
	keys := func(token *stdjwt.Token) (interface{}, error) {
		return key, nil
	}

	jwtOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(auth.AuthErrorEncoder),
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(jwt.ToHTTPContext()),
	}



	//jwt


	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "user_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "user_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "user_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here


	var user usersvc.UserService
	user = usersvc.Userservice{}
	user = usersvc.LoggingMiddleware{logger, user}
	user = usersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, user}

	var school  schoolsvc.SchoolService
	school = schoolsvc.Schoolservice{}
	school = schoolsvc.LoggingMiddleware{*loggrer,school}
	school = schoolsvc.InstrumentingMiddleware{requestCount,requestLatency,countResult,school}


	var userCreateEndpoint endpoint.Endpoint
	{
		userCreateEndpoint = usersvc.MakeCreateEndpoint(user)
		userCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(userCreateEndpoint)
	}
	var getOneUserEndpoint endpoint.Endpoint
	{
		getOneUserEndpoint = usersvc.MakeGetOneEndpoint(user)
		getOneUserEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getOneUserEndpoint)
	}
	var getAllUsersEndpoint endpoint.Endpoint
	{
		getAllUsersEndpoint = usersvc.MakeGetAllEndpoint(user)
		getAllUsersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getAllUsersEndpoint)
	}

//school endpoint
	var schoolCreateEndpoint endpoint.Endpoint
	{
		schoolCreateEndpoint = schoolsvc.MakeCreateEndpoint(school)
		schoolCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(schoolCreateEndpoint)
	}
	var getOneSchoolEndpoint endpoint.Endpoint
	{
		getOneSchoolEndpoint = schoolsvc.MakeGetOneEndpoint(school)
		getOneSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getOneSchoolEndpoint)
	}
	var getAllSchoolsEndpoint endpoint.Endpoint
	{
		getAllSchoolsEndpoint = schoolsvc.MakeGetAllEndpoint(school)
		getAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getAllSchoolsEndpoint)
	}

	var getBestSchoolEndpoint endpoint.Endpoint
	{
		getBestSchoolEndpoint = schoolsvc.MakeRetrieveBestPerfomingSchoolEndpoint(school)
		getBestSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getBestSchoolEndpoint)
	}

	var recordSchoolPerformanceEndpoint endpoint.Endpoint
	{
		recordSchoolPerformanceEndpoint = schoolsvc.MakeRecordPerformanceEndpoint(school)
		recordSchoolPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(recordSchoolPerformanceEndpoint)
	}
	var rankAllSchoolsEndpoint endpoint.Endpoint
	{
		rankAllSchoolsEndpoint = schoolsvc.MakeRankAllSchoolsEndpoint(school)
		rankAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(rankAllSchoolsEndpoint)
	}



	authFieldKeys := []string{"method", "error"}
	requestAuthCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "auth_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, authFieldKeys)
	requestAuthLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "auth_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, authFieldKeys)

	// API clients database
	var clients = map[string]string{
		"mobile": "m_secret",
		"web":    "w_secret",
	}
	var authy auth.AuthService
	authy = auth.Authservice{key, clients}
	authy = auth.LoggingAuthMiddleware{logger, authy.(auth.Authservice)}
	authy = auth.InstrumentingAuthMiddleware{requestAuthCount, requestAuthLatency, authy}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(auth.AuthErrorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	authHandler := httptransport.NewServer(
		auth.MakeAuthEndpoint(authy),
		auth.DecodeAuthRequest,
		auth.EncodeResponse,
		options...,
	)
	//
	userHandler := httptransport.NewServer(
		userCreateEndpoint,
		usersvc.DecodeCreateRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)
	getOneUserHandler := httptransport.NewServer(
		getOneUserEndpoint,
		usersvc.DecodeGetOneRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)
	getAllUsersHandler := httptransport.NewServer(
		getAllUsersEndpoint,
		usersvc.DecodeGetAllRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)

	schoolHandler := httptransport.NewServer(
		schoolCreateEndpoint,
		schoolsvc.DecodeCreateRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)

	recordPerformanceHandler := httptransport.NewServer(
		recordSchoolPerformanceEndpoint,
		schoolsvc.DecodeRecordPerformanceRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneSchoolHandler := httptransport.NewServer(
		getOneSchoolEndpoint,
		schoolsvc.DecodeGetOneRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)

	getBestSchoolHandler := httptransport.NewServer(
		getBestSchoolEndpoint,
		schoolsvc.DecodeGetBestSchoolRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllSchoolsHandler := httptransport.NewServer(
		getAllSchoolsEndpoint,
		schoolsvc.DecodeGetAllRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)
	rankAllSchoolsHandler := httptransport.NewServer(
		rankAllSchoolsEndpoint,
		schoolsvc.DecodeRankAllSchoolsRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)


	var routes = Routes{
		Route{
			"User",
			"POST",
			"/user",
			userHandler,
		},
		Route{
			"User ",
			"GET",
			"/user/{id}",
			getOneUserHandler,
		},
		Route{
			"Users ",
			"GET",
			"/users",
			getAllUsersHandler,
		},
		Route{
			"School",
			"POST",
			"/school",
			schoolHandler,
		},
		Route{
			"School Performance",
			"POST",
			"/school_performance",
			recordPerformanceHandler,
		},
		Route{
			"School ",
			"GET",
			"/school/{id}",
			getOneSchoolHandler,
		},
		Route{
			"Best School ",
			"POST",
			"/best_school",
			getBestSchoolHandler,
		},
		Route{
			"Schools ",
			"GET",
			"/schools",
			getAllSchoolsHandler,
		},
		Route{
			"Rank Schools ",
			"POST",
			"/ranking",
			rankAllSchoolsHandler,
		},

		Route{
			"Auth",
			"POST",
			"/auth",
			authHandler,
		},

	}
	r := APINewRouter(routes)
	version1 := r.PathPrefix("/v1").Subrouter()
	version2 := r.PathPrefix("/v2").Subrouter()
	AddRoutes(version1,routes)
	AddRoutes(version2,routes)
	//r.Handle()
	r.Handle("/metrics", stdprometheus.Handler())
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", r,))
	}

	//
	//logger.Log("msg", "HTTP", "addr", ":8080")
	//logger.Log("err", http.ListenAndServe(":8080", nil))



//order {"a":1,"b":4}
//customer {"s":"This is it"}
//partner {"s":"I am a partner"}

func basicAuth(username string, password string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised\n"))
			return
		}

		h.ServeHTTP(w, r)
	})
}