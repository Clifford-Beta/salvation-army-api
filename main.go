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
	catsvc "salv_prj/category"
	"salv_prj/auth"
	"salv_prj/store"
	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/rs/cors"

)

var logger = logrus.New()

func main() {

	//store.ConfigureApp("staging")
	defer store.Database.Close()

	//ctx := context.Background()
	logit := log.NewLogfmtLogger(os.Stderr)
	logger.Out = os.Stdout

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
		httptransport.ServerErrorLogger(logit),
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
	user = usersvc.LoggingMiddleware{*logger, user}
	user = usersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, user}

	var school  schoolsvc.SchoolService
	school = schoolsvc.Schoolservice{}
	school = schoolsvc.LoggingMiddleware{*logger,school}
	school = schoolsvc.InstrumentingMiddleware{requestCount,requestLatency,countResult,school}

	var category  catsvc.CategoryService
	category = catsvc.Categoryservice{}
	category = catsvc.LoggingMiddleware{*logger,category}
	category = catsvc.InstrumentingMiddleware{requestCount,requestLatency,countResult,category}



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

	var loginUserEndpoint endpoint.Endpoint
	{
		loginUserEndpoint = usersvc.MakeLoginEndpoint(user)
		loginUserEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(loginUserEndpoint)
	}
	var getAllUsersEndpoint endpoint.Endpoint
	{
		getAllUsersEndpoint = usersvc.MakeGetAllEndpoint(user)
		getAllUsersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getAllUsersEndpoint)
	}

	//category endpoint

	var categoryCreateEndpoint endpoint.Endpoint
	{
		categoryCreateEndpoint = catsvc.MakeCreateEndpoint(category)
		categoryCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(categoryCreateEndpoint)
	}
	var getOneCategoryEndpoint endpoint.Endpoint
	{
		getOneCategoryEndpoint = catsvc.MakeGetOneEndpoint(category)
		getOneCategoryEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getOneCategoryEndpoint)
	}
	var getAllCategoriesEndpoint endpoint.Endpoint
	{
		getAllCategoriesEndpoint = catsvc.MakeGetAllEndpoint(category)
		getAllCategoriesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getAllCategoriesEndpoint)
	}
	var tierCreateEndpoint endpoint.Endpoint
	{
		tierCreateEndpoint = catsvc.MakeCreateTierEndpoint(category)
		tierCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(tierCreateEndpoint)
	}
	var getOneTierEndpoint endpoint.Endpoint
	{
		getOneTierEndpoint = catsvc.MakeGetOneTierEndpoint(category)
		getOneTierEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getOneTierEndpoint)
	}
	var getAllTiersEndpoint endpoint.Endpoint
	{
		getAllTiersEndpoint = catsvc.MakeGetAllTiersEndpoint(category)
		getAllTiersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(getAllTiersEndpoint)
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
	var clients = map[string]int{
		"mobile": 1,
		"web":    2,
	}
	var authy auth.AuthService
	authy = auth.Authservice{key, clients}
	authy = auth.LoggingAuthMiddleware{*logger, authy.(auth.Authservice)}
	authy = auth.InstrumentingAuthMiddleware{requestAuthCount, requestAuthLatency, authy}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(auth.AuthErrorEncoder),
		httptransport.ServerErrorLogger(logit),
	}

	authHandler := httptransport.NewServer(
		auth.MakeAuthEndpoint(authy),
		auth.DecodeAuthRequest,
		auth.EncodeResponse,
		options...,
	)
	// user handelers
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

	loginUserHandler := httptransport.NewServer(
		loginUserEndpoint,
		usersvc.DecodeLoginRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)
	getAllUsersHandler := httptransport.NewServer(
		getAllUsersEndpoint,
		usersvc.DecodeGetAllRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)

	//category handlers

	createCategoryHandler := httptransport.NewServer(
		categoryCreateEndpoint,
		catsvc.DecodeCreateRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)
	createTierHandler := httptransport.NewServer(
		tierCreateEndpoint,
		catsvc.DecodeCreateTierRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneCategoryHandler := httptransport.NewServer(
		getOneCategoryEndpoint,
		catsvc.DecodeGetOneRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneTierHandler := httptransport.NewServer(
		getOneTierEndpoint,
		catsvc.DecodeGetOneRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllCategoriesHandler := httptransport.NewServer(
		getAllCategoriesEndpoint,
		catsvc.DecodeGetAllRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllTiersHandler := httptransport.NewServer(
		getAllTiersEndpoint,
		catsvc.DecodeGetAllRequest,
		catsvc.EncodeResponse,
		jwtOptions...,
	)

	//school handlers

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
			"/user",
			getAllUsersHandler,
		},
		Route{
			"Login ",
			"POST",
			"/login",
			loginUserHandler,
		},
		Route{
			"Category",
			"POST",
			"/category",
			createCategoryHandler,
		},
		Route{
			"Category ",
			"GET",
			"/category/{id}",
			getOneCategoryHandler,
		},
		Route{
			"categories ",
			"GET",
			"/category",
			getAllCategoriesHandler,
		},
		Route{
			"Tier",
			"POST",
			"/tier",
			createTierHandler,
		},
		Route{
			"Tier ",
			"GET",
			"/tier/{id}",
			getOneTierHandler,
		},
		Route{
			"categories ",
			"GET",
			"/tier",
			getAllTiersHandler,
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
			"POST",
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
	handler := cors.Default().Handler(r)
	version1 := r.PathPrefix("/v1").Subrouter()
	version2 := r.PathPrefix("/v2").Subrouter()
	AddRoutes(version1,routes)
	AddRoutes(version2,routes)
	//r.Handle()
	r.Handle("/metrics", stdprometheus.Handler())
	logger.WithFields(logrus.Fields{"msg": "HTTP", "addr": ":8000"}).Info("Everything is ready, let's go !!!")
	logger.WithFields(logrus.Fields{"err": http.ListenAndServe(":8000", corsHandler(handler))}).Fatal("Oops! the server crashed")
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


func corsHandler(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.WriteHeader(204)
			logrus.Debug("I got here")
			return
		}
		h.ServeHTTP(w,r)
	})
}