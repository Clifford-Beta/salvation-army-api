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
	"salv_prj/auth"
	"salv_prj/store"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/endpoint"
)


func main() {

	//store.ConfigureApp("staging")
	defer store.Database.Close()

	//ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)
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

	//var customer customersvc.CustomerService
	//customer = customersvc.Customeservice{}
	//customer = customersvc.LoggingMiddleware{logger, customer}
	//customer = customersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, customer}
	//
	//
	//var order ordersvc.OrderService
	//order = ordersvc.Orderservice{}
	//order = ordersvc.LoggingMiddleware{logger, order}
	//order = ordersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, order}
	//
	//var partner partnersvc.PartnerService
	//partner = partnersvc.Partnerservice{}
	//partner = partnersvc.LoggingMiddleware{logger, partner}
	//partner = partnersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, partner}


	var user usersvc.UserService
	user = usersvc.Userservice{}
	user = usersvc.LoggingMiddleware{logger, user}
	user = usersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, user}




	//var customerTestEndpoint endpoint.Endpoint
	//{
	//	customerTestEndpoint = customersvc.MakeTestEndpoint(customer)
	//	customerTestEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(customerTestEndpoint)
	//	//customerTestEndpointcustomerTestEndpoint = jwt.NewSigner(
	//	//	"kid",
	//	//	[]byte("SigningString"),
	//	//	stdjwt.SigningMethodHS256,
	//	//	jwt.Claims{"user":"Beta C W"},
	//	//)(customerTestEndpoint)
	//}
	//
	//
	//var ordersTestEndpoint endpoint.Endpoint
	//{
	//	//kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("SigningString"), nil }
	//	ordersTestEndpoint = ordersvc.MakeAddEndpoint(order)
	//	ordersTestEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(ordersTestEndpoint)
	//}
	//
	//var partnerTestEndpoint endpoint.Endpoint
	//{
	//	partnerTestEndpoint = partnersvc.MakeTestEndpoint(partner)
	//	partnerTestEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,&auth.CustomClaims{})(partnerTestEndpoint)
	//}


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
//auth endpoint



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


	var routes = Routes{

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
		//Route{
		//	"Order Test",
		//	"POST",
		//	"/order",
		//	orderTestHandler,
		//},
		//Route{
		//	"Partner Test",
		//	"POST",
		//	"/partner",
		//	partnerTestHandler,
		//},
		Route{
			"Auth",
			"POST",
			"/auth",
			authHandler,
		},
		Route{
			"User",
			"POST",
			"/user",
			userHandler,
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

	//http.Handle("/uppercase", uppercaseHandler)
	//http.Handle("/count", countHandler)
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