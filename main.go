package main

import (
	"crypto/subtle"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	//customersvc "salvation-army-api/customerservice"
	//ordersvc "salvation-army-api/orderservice"
	//partnersvc "salvation-army-api/partnerservice"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	actsvc "salvation-army-api/activity"
	"salvation-army-api/auth"
	catsvc "salvation-army-api/category"
	filesvc "salvation-army-api/file"
	infsvc "salvation-army-api/infrastructure"
	msgsvc "salvation-army-api/message"
	projectsvc "salvation-army-api/project"
	schoolsvc "salvation-army-api/school"
	staffsvc "salvation-army-api/staff"
	"salvation-army-api/store"
	usersvc "salvation-army-api/user"
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
		httptransport.ServerBefore(jwt.HTTPToContext()),
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

	//file service
	var file filesvc.FileService
	file = filesvc.Fileservice{}
	file = filesvc.LoggingMiddleware{*logger, file}
	//user = usersvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, user}
	//message service
	var message msgsvc.MessageService
	message = msgsvc.Messsageservice{}
	message = msgsvc.LoggingMiddleware{*logger, message}

	//project service
	var project projectsvc.ProjectService
	project = projectsvc.Projectservice{}
	project = projectsvc.LoggingMiddleware{*logger, project}

	//infrastructure service
	var infs infsvc.InfrastructureService
	infs = infsvc.Infrastructureservice{}
	infs = infsvc.LoggingMiddleware{*logger, infs}

	//activity service

	var activity actsvc.ActivityService
	activity = actsvc.Activityservice{}
	activity = actsvc.LoggingMiddleware{*logger, activity}

	//staff service
	var staff staffsvc.StaffService
	staff = staffsvc.Staffservice{}
	staff = staffsvc.LoggingMiddleware{*logger, staff}

	//school service
	var school schoolsvc.SchoolService
	school = schoolsvc.Schoolservice{}
	school = schoolsvc.LoggingMiddleware{*logger, school}
	school = schoolsvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, school}

	var category catsvc.CategoryService
	category = catsvc.Categoryservice{}
	category = catsvc.LoggingMiddleware{*logger, category}
	category = catsvc.InstrumentingMiddleware{requestCount, requestLatency, countResult, category}

	//user endpoints
	var userCreateEndpoint endpoint.Endpoint
	{
		userCreateEndpoint = usersvc.MakeCreateEndpoint(user)
		userCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory,)(userCreateEndpoint)
	}

	var userUpdateEndpoint endpoint.Endpoint
	{
		userUpdateEndpoint = usersvc.MakeUpdateteEndpoint(user)
		userUpdateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory,)(userUpdateEndpoint)
	}

	var userDeleteEndpoint endpoint.Endpoint
	{
		userUpdateEndpoint = usersvc.MakeDeleteEndpoint(user)
		userDeleteEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory,)(userDeleteEndpoint)
	}
	var getOneUserEndpoint endpoint.Endpoint
	{
		getOneUserEndpoint = usersvc.MakeGetOneEndpoint(user)
		getOneUserEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneUserEndpoint)
	}


	var getAllUsersEndpoint endpoint.Endpoint
	{
		getAllUsersEndpoint = usersvc.MakeGetAllEndpoint(user)
		getAllUsersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllUsersEndpoint)
	}
	// user handelers
	userHandler := httptransport.NewServer(

		userCreateEndpoint,
		usersvc.DecodeCreateRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)

	updateUserHandler := httptransport.NewServer(
		userUpdateEndpoint,
		usersvc.DecodeCreateRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)

	deleteUserHandler := httptransport.NewServer(
		userDeleteEndpoint,
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

	//loginUserHandler := httptransport.NewServer(
	//
	//	loginUserEndpoint,
	//	usersvc.DecodeLoginRequest,
	//	usersvc.EncodeResponse,
	//	jwtOptions...,
	//)
	getAllUsersHandler := httptransport.NewServer(

		getAllUsersEndpoint,
		usersvc.DecodeGetAllRequest,
		usersvc.EncodeResponse,
		jwtOptions...,
	)

	//category endpoint

	var categoryCreateEndpoint endpoint.Endpoint
	{
		categoryCreateEndpoint = catsvc.MakeCreateEndpoint(category)
		categoryCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(categoryCreateEndpoint)
	}
	var getOneCategoryEndpoint endpoint.Endpoint
	{
		getOneCategoryEndpoint = catsvc.MakeGetOneEndpoint(category)
		getOneCategoryEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneCategoryEndpoint)
	}
	var getAllCategoriesEndpoint endpoint.Endpoint
	{
		getAllCategoriesEndpoint = catsvc.MakeGetAllEndpoint(category)
		getAllCategoriesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllCategoriesEndpoint)
	}
	var tierCreateEndpoint endpoint.Endpoint
	{
		tierCreateEndpoint = catsvc.MakeCreateTierEndpoint(category)
		tierCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(tierCreateEndpoint)
	}
	var getOneTierEndpoint endpoint.Endpoint
	{
		getOneTierEndpoint = catsvc.MakeGetOneTierEndpoint(category)
		getOneTierEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneTierEndpoint)
	}
	var getAllTiersEndpoint endpoint.Endpoint
	{
		getAllTiersEndpoint = catsvc.MakeGetAllTiersEndpoint(category)
		getAllTiersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllTiersEndpoint)
	}

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
	//school endpoint
	var schoolCreateEndpoint endpoint.Endpoint
	{
		schoolCreateEndpoint = schoolsvc.MakeCreateEndpoint(school)
		schoolCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(schoolCreateEndpoint)
	}

	var schoolUpdateEndpoint endpoint.Endpoint
	{
		schoolUpdateEndpoint = schoolsvc.MakeUpdateEndpoint(school)
		schoolUpdateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(schoolUpdateEndpoint)
	}

	var schoolDeleteEndpoint endpoint.Endpoint
	{
		schoolDeleteEndpoint = schoolsvc.MakeDeleteEndpoint(school)
		schoolDeleteEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(schoolDeleteEndpoint)
	}
	var getOneSchoolEndpoint endpoint.Endpoint
	{
		getOneSchoolEndpoint = schoolsvc.MakeGetOneEndpoint(school)
		getOneSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneSchoolEndpoint)
	}
	var getAllSchoolsEndpoint endpoint.Endpoint
	{
		getAllSchoolsEndpoint = schoolsvc.MakeGetAllEndpoint(school)
		getAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllSchoolsEndpoint)
	}

	var getBestSchoolEndpoint endpoint.Endpoint
	{
		getBestSchoolEndpoint = schoolsvc.MakeRetrieveBestPerfomingSchoolEndpoint(school)
		getBestSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getBestSchoolEndpoint)
	}

	var recordSchoolPerformanceEndpoint endpoint.Endpoint
	{
		recordSchoolPerformanceEndpoint = schoolsvc.MakeRecordPerformanceEndpoint(school)
		recordSchoolPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(recordSchoolPerformanceEndpoint)
	}
	var rankAllSchoolsEndpoint endpoint.Endpoint
	{
		rankAllSchoolsEndpoint = schoolsvc.MakeRankAllSchoolsEndpoint(school)
		rankAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(rankAllSchoolsEndpoint)
	}

	//school handlers

	schoolHandler := httptransport.NewServer(

		schoolCreateEndpoint,
		schoolsvc.DecodeCreateRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)

	updateSchoolHandler := httptransport.NewServer(
		schoolUpdateEndpoint,
		schoolsvc.DecodeCreateRequest,
		schoolsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteSchoolHandler := httptransport.NewServer(
		schoolDeleteEndpoint,
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

	//activity endpoint
	var createActivityEndpoint endpoint.Endpoint
	{
		createActivityEndpoint = actsvc.MakeCreateActivityEndpoint(activity)
		createActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createActivityEndpoint)
	}

	var updateActivityEndpoint endpoint.Endpoint
	{
		updateActivityEndpoint = actsvc.MakeUpdateActivityEndpoint(activity)
		updateActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateActivityEndpoint)
	}

	var deleteActivityEndpoint endpoint.Endpoint
	{
		deleteActivityEndpoint = actsvc.MakeDeleteActivityEndpoint(activity)
		deleteActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteActivityEndpoint)
	}
	var createLevelEndpoint endpoint.Endpoint
	{
		createLevelEndpoint = actsvc.MakeCreateLevelEndpoint(activity)
		createLevelEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createLevelEndpoint)
	}
	var recordPerformanceEndpoint endpoint.Endpoint
	{
		recordPerformanceEndpoint = actsvc.MakeRecordPerformanceEndpoint(activity)
		recordPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(recordPerformanceEndpoint)
	}
	var getOneActivityEndpoint endpoint.Endpoint
	{
		getOneActivityEndpoint = actsvc.MakeGetOneActivityEndpoint(activity)
		getOneActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneActivityEndpoint)
	}

	var getOneLevelEndpoint endpoint.Endpoint
	{
		getOneLevelEndpoint = actsvc.MakeGetOneLevelEndpoint(activity)
		getOneLevelEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneLevelEndpoint)
	}
	var getOnePerformanceEndpoint endpoint.Endpoint
	{
		getOnePerformanceEndpoint = actsvc.MakeGetOnePerformanceEndpoint(activity)
		getOnePerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOnePerformanceEndpoint)
	}
	var getAllActivitiesEndpoint endpoint.Endpoint
	{
		getAllActivitiesEndpoint = actsvc.MakeGetAllActivitiesEndpoint(activity)
		getAllActivitiesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllActivitiesEndpoint)
	}
	var getAllLevelsEndpoint endpoint.Endpoint
	{
		getAllLevelsEndpoint = actsvc.MakeGetAllLevelsEndpoint(activity)
		getAllLevelsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllLevelsEndpoint)
	}
	var getAllPerformancesEndpoint endpoint.Endpoint
	{
		getAllPerformancesEndpoint = actsvc.MakeGetAllPerformancesEndpoint(activity)
		getAllPerformancesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllPerformancesEndpoint)
	}

	//activity handlers
	createActivityHandler := httptransport.NewServer(

		createActivityEndpoint,
		actsvc.DecodeCreateActivityRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)

	updateActivityHandler := httptransport.NewServer(
		updateActivityEndpoint,
		actsvc.DecodeCreateActivityRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteActivityHandler := httptransport.NewServer(
		deleteActivityEndpoint,
		actsvc.DecodeCreateActivityRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	createLevelHandler := httptransport.NewServer(

		createLevelEndpoint,
		actsvc.DecodeCreateLevelRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	recordactivityPerformanceHandler := httptransport.NewServer(

		recordPerformanceEndpoint,
		actsvc.DecodeRecordPerformanceRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneActivityHandler := httptransport.NewServer(

		getOneActivityEndpoint,
		actsvc.DecodeGetOneRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneLevelHandler := httptransport.NewServer(

		getOneLevelEndpoint,
		actsvc.DecodeGetOneRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getOnePerformanceHandler := httptransport.NewServer(

		getOnePerformanceEndpoint,
		actsvc.DecodeGetOneRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllActivitiesHandler := httptransport.NewServer(

		getAllActivitiesEndpoint,
		actsvc.DecodeGetAllRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllLevelsHandler := httptransport.NewServer(

		getAllLevelsEndpoint,
		actsvc.DecodeGetAllRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllPerformancesHandler := httptransport.NewServer(

		getAllPerformancesEndpoint,
		actsvc.DecodeGetAllRequest,
		actsvc.EncodeResponse,
		jwtOptions...,
	)

	//file endpoints
	var createFileEndpoint endpoint.Endpoint
	{
		createFileEndpoint = filesvc.MakeCreateEndpoint(file)
		createFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createFileEndpoint)
	}

	var updateFileEndpoint endpoint.Endpoint
	{
		updateFileEndpoint = filesvc.MakeUpdateEndpoint(file)
		updateFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateFileEndpoint)
	}

	var deleteFileEndpoint endpoint.Endpoint
	{
		deleteFileEndpoint = filesvc.MakeDeleteEndpoint(file)
		deleteFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteFileEndpoint)
	}

	var createFileTypeEndpoint endpoint.Endpoint
	{
		createFileTypeEndpoint = filesvc.MakeCreateTypeEndpoint(file)
		createFileTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createFileTypeEndpoint)
	}

	var getOneFileEndpoint endpoint.Endpoint
	{
		getOneFileEndpoint = filesvc.MakeGetOneEndpoint(file)
		getOneFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneFileEndpoint)
	}

	var getOneFileTypeEndpoint endpoint.Endpoint
	{
		getOneFileTypeEndpoint = filesvc.MakeGetOneTypeEndpoint(file)
		getOneFileTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneFileTypeEndpoint)
	}

	var getAllFilesEndpoint endpoint.Endpoint
	{
		getAllFilesEndpoint = filesvc.MakeGetAllEndpoint(file)
		getAllFilesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllFilesEndpoint)
	}
	var getAllFileTypesEndpoint endpoint.Endpoint
	{
		getAllFileTypesEndpoint = filesvc.MakeGetAllTypesEndpoint(file)
		getAllFileTypesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllFileTypesEndpoint)
	}

	//file handlers

	createFileHandler := httptransport.NewServer(

		createFileEndpoint,
		filesvc.DecodeCreateRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	updateFileHandler := httptransport.NewServer(

		updateFileEndpoint,
		filesvc.DecodeCreateRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	deleteFileHandler := httptransport.NewServer(

		deleteFileEndpoint,
		filesvc.DecodeCreateRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	createFileTypeHandler := httptransport.NewServer(

		createFileTypeEndpoint,
		filesvc.DecodeCreateTypeRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	getOneFileHandler := httptransport.NewServer(

		getOneFileEndpoint,
		filesvc.DecodeGetOneRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	getOneFileTypeHandler := httptransport.NewServer(

		getOneFileTypeEndpoint,
		filesvc.DecodeGetOneRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	getAllFilesHandler := httptransport.NewServer(

		getAllFilesEndpoint,
		filesvc.DecodeGetAllRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)
	getAllFileTypesHandler := httptransport.NewServer(

		getAllFileTypesEndpoint,
		filesvc.DecodeGetAllRequest,
		filesvc.EncodeResponse,
		jwtOptions...,
	)

	//insfrastructure endpoints
	var createInfEndpoint endpoint.Endpoint
	{
		createInfEndpoint = infsvc.MakeCreateEndpoint(infs)
		createInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createInfEndpoint)
	}

	var updateInfEndpoint endpoint.Endpoint
	{
		updateInfEndpoint = infsvc.MakeUpdateEndpoint(infs)
		updateInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateInfEndpoint)
	}

	var deleteInfEndpoint endpoint.Endpoint
	{
		deleteInfEndpoint = infsvc.MakeDeleteEndpoint(infs)
		deleteInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteInfEndpoint)
	}

	var createInfTypeEndpoint endpoint.Endpoint
	{
		createInfTypeEndpoint = infsvc.MakeCreateTypeEndpoint(infs)
		createInfTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createInfTypeEndpoint)
	}

	var getOneInfEndpoint endpoint.Endpoint
	{
		getOneInfEndpoint = infsvc.MakeGetOneEndpoint(infs)
		getOneInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneInfEndpoint)
	}

	var getOneInfTypeEndpoint endpoint.Endpoint
	{
		getOneInfTypeEndpoint = infsvc.MakeGetOneTypeEndpoint(infs)
		getOneInfTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneInfTypeEndpoint)
	}

	var getAllInfEndpoint endpoint.Endpoint
	{
		getAllInfEndpoint = infsvc.MakeGetAllEndpoint(infs)
		getAllInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllInfEndpoint)
	}
	var getAllInfTypesEndpoint endpoint.Endpoint
	{
		getAllInfTypesEndpoint = infsvc.MakeGetAllTypesEndpoint(infs)
		getAllInfTypesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllInfTypesEndpoint)
	}

	//infrastructure handlers
	createInfHandler := httptransport.NewServer(

		createInfEndpoint,
		infsvc.DecodeCreateRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	updateInfHandler := httptransport.NewServer(

		updateInfEndpoint,
		infsvc.DecodeCreateRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteInfHandler := httptransport.NewServer(

		deleteInfEndpoint,
		infsvc.DecodeCreateRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	createInfTypeHandler := httptransport.NewServer(

		createInfTypeEndpoint,
		infsvc.DecodeCreateTypeRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneInfHandler := httptransport.NewServer(
		getOneInfEndpoint,
		infsvc.DecodeGetOneRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	getOneInfTypeHandler := httptransport.NewServer(

		getOneInfTypeEndpoint,
		infsvc.DecodeGetOneRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	getAllInfHandler := httptransport.NewServer(

		getAllInfEndpoint,
		infsvc.DecodeGetAllRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllInfTypesHandler := httptransport.NewServer(

		getAllInfTypesEndpoint,
		infsvc.DecodeGetAllRequest,
		infsvc.EncodeResponse,
		jwtOptions...,
	)

	//message endpoints
	var createMessageEndpoint endpoint.Endpoint
	{
		createMessageEndpoint = msgsvc.MakeCreateEndpoint(message)
		createMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createMessageEndpoint)
	}

	var updateMessageEndpoint endpoint.Endpoint
	{
		updateMessageEndpoint = msgsvc.MakeUpdateEndpoint(message)
		updateMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateMessageEndpoint)
	}

	var deleteMessageEndpoint endpoint.Endpoint
	{
		deleteMessageEndpoint = msgsvc.MakeDeleteEndpoint(message)
		deleteMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteMessageEndpoint)
	}
	var getOneMessageEndpoint endpoint.Endpoint
	{
		getOneMessageEndpoint = msgsvc.MakeGetOneEndpoint(message)
		getOneMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneMessageEndpoint)
	}
	var getAllMessagesEndpoint endpoint.Endpoint
	{
		getAllMessagesEndpoint = msgsvc.MakeGetAllEndpoint(message)
		getAllMessagesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllMessagesEndpoint)
	}

	//message handlers
	createMessageHandler := httptransport.NewServer(

		createMessageEndpoint,
		msgsvc.DecodeCreateRequest,
		msgsvc.EncodeResponse,
		jwtOptions...,
	)

	updateMessageHandler := httptransport.NewServer(

		updateMessageEndpoint,
		msgsvc.DecodeCreateRequest,
		msgsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteMessageHandler := httptransport.NewServer(

		deleteMessageEndpoint,
		msgsvc.DecodeCreateRequest,
		msgsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneMessageHandler := httptransport.NewServer(

		getOneMessageEndpoint,
		msgsvc.DecodeGetOneRequest,
		msgsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllMessagesHandler := httptransport.NewServer(

		getAllMessagesEndpoint,
		msgsvc.DecodeGetAllRequest,
		msgsvc.EncodeResponse,
		jwtOptions...,
	)

	//project endpoints
	var createProjectEndpoint endpoint.Endpoint
	{
		createProjectEndpoint = projectsvc.MakeCreateEndpoint(project)
		createProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(createProjectEndpoint)
	}

	var updateProjectEndpoint endpoint.Endpoint
	{
		updateProjectEndpoint = projectsvc.MakeUpdateEndpoint(project)
		updateProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateProjectEndpoint)
	}

	var deleteProjectEndpoint endpoint.Endpoint
	{
		deleteProjectEndpoint = projectsvc.MakeDeleteEndpoint(project)
		deleteProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteProjectEndpoint)
	}
	var getOneProjectEndpoint endpoint.Endpoint
	{
		getOneProjectEndpoint = projectsvc.MakeGetOneEndpoint(project)
		getOneProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getOneProjectEndpoint)
	}
	var getAllProjectEndpoint endpoint.Endpoint
	{
		getAllProjectEndpoint = projectsvc.MakeGetAllEndpoint(project)
		getAllProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(getAllProjectEndpoint)
	}

	//project handler

	createProjectHandler := httptransport.NewServer(

		createProjectEndpoint,
		projectsvc.DecodeCreateRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)

	updateProjectHandler := httptransport.NewServer(

		updateProjectEndpoint,
		projectsvc.DecodeCreateRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteProjectHandler := httptransport.NewServer(

		deleteProjectEndpoint,
		projectsvc.DecodeCreateRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneProjectHandler := httptransport.NewServer(

		getOneProjectEndpoint,
		projectsvc.DecodeGetOneRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)
	getAllProjectHandler := httptransport.NewServer(

		getAllProjectEndpoint,
		projectsvc.DecodeGetAllRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)
	// staff endpoints
	var addStaffEndpoint endpoint.Endpoint
	{
		addStaffEndpoint = staffsvc.MakeAddStaffEndpoint(staff)
		addStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(addStaffEndpoint)
	}

	var updateStaffEndpoint endpoint.Endpoint
	{
		updateStaffEndpoint = staffsvc.MakeUpdateStaffEndpoint(staff)
		updateStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateStaffEndpoint)
	}

	var deleteStaffEndpoint endpoint.Endpoint
	{
		deleteStaffEndpoint = staffsvc.MakeDeleteStaffEndpoint(staff)
		deleteStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteStaffEndpoint)
	}
	var retrieveStaffEndpoint endpoint.Endpoint
	{
		retrieveStaffEndpoint = staffsvc.MakeRetrieveStaffEndpoint(staff)
		retrieveStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveStaffEndpoint)
	}
	var retrieveAllStaffEndpoint endpoint.Endpoint
	{
		retrieveAllStaffEndpoint = staffsvc.MakeRetrieveAllStaffEndpoint(staff)
		retrieveAllStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveAllStaffEndpoint)
	}
	var addRoleEndpoint endpoint.Endpoint
	{
		addRoleEndpoint = staffsvc.MakeAddStaffRoleEndpoint(staff)
		addRoleEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(addRoleEndpoint)
	}
	var retrieveRoleEndpoint endpoint.Endpoint
	{
		retrieveRoleEndpoint = staffsvc.MakeRetrieveStaffRoleEndpoint(staff)
		retrieveRoleEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveRoleEndpoint)
	}
	var recordBestPerformingStaffEndpoint endpoint.Endpoint
	{
		recordBestPerformingStaffEndpoint = staffsvc.MakeRecordBestPerformingStaffEndpoint(staff)
		recordBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(recordBestPerformingStaffEndpoint)
	}

	var updateBestPerformingStaffEndpoint endpoint.Endpoint
	{
		updateBestPerformingStaffEndpoint = staffsvc.MakeUpdateBestPerformingStaffEndpoint(staff)
		updateBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updateBestPerformingStaffEndpoint)
	}

	var deleteBestPerformingStaffEndpoint endpoint.Endpoint
	{
		deleteBestPerformingStaffEndpoint = staffsvc.MakeDeleteBestPerformingStaffEndpoint(staff)
		deleteBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deleteBestPerformingStaffEndpoint)
	}

	var bestPerformingStudentEndpoint endpoint.Endpoint
	{
		bestPerformingStudentEndpoint = staffsvc.MakeRecordBestPerformingStudentEndpoint(staff)
		bestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(bestPerformingStudentEndpoint)
	}

	var updatebestPerformingStudentEndpoint endpoint.Endpoint
	{
		updatebestPerformingStudentEndpoint = staffsvc.MakeUpdateBestPerformingStudentEndpoint(staff)
		updatebestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(updatebestPerformingStudentEndpoint)
	}

	var deletebestPerformingStudentEndpoint endpoint.Endpoint
	{
		deletebestPerformingStudentEndpoint = staffsvc.MakeDeleteBestPerformingStudentEndpoint(staff)
		deletebestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(deletebestPerformingStudentEndpoint)
	}
	var retrieveTeacherEndpoint endpoint.Endpoint
	{
		retrieveTeacherEndpoint = staffsvc.MakeRetrieveTeacherEndpoint(staff)
		retrieveTeacherEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveTeacherEndpoint)
	}

	var retrieveStudentEndpoint endpoint.Endpoint
	{
		retrieveStudentEndpoint = staffsvc.MakeRetrieveStudentEndpoint(staff)
		retrieveStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveStudentEndpoint)
	}
	var retrieveBestPerformingStaffEndpoint endpoint.Endpoint
	{
		retrieveBestPerformingStaffEndpoint = staffsvc.MakeRetrieveBestPerformingStaffEndpoint(staff)
		retrieveBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveBestPerformingStaffEndpoint)
	}
	var retrieveBestPerformingStudentEndpoint endpoint.Endpoint
	{
		retrieveBestPerformingStudentEndpoint = staffsvc.MakeRetrieveBestPerformingStudentEndpoint(staff)
		retrieveBestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(retrieveBestPerformingStudentEndpoint)
	}
	var rankStaffPerformanceEndpoint endpoint.Endpoint
	{
		rankStaffPerformanceEndpoint = staffsvc.MakeRankStaffPerformanceEndpoint(staff)
		rankStaffPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(rankStaffPerformanceEndpoint)
	}
	var rankStudentPerformanceEndpoint endpoint.Endpoint
	{
		rankStudentPerformanceEndpoint = staffsvc.MakeRankStudentPerformanceEndpoint(staff)
		rankStudentPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256,auth.MapClaimsFactory)(rankStudentPerformanceEndpoint)
	}

	//staff handlers

	addStaffHandler := httptransport.NewServer(

		addStaffEndpoint,
		staffsvc.DecodeAddStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	updateStaffHandler := httptransport.NewServer(

		updateStaffEndpoint,
		staffsvc.DecodeAddStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteStaffHandler := httptransport.NewServer(

		deleteStaffEndpoint,
		staffsvc.DecodeAddStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveStaffHandler := httptransport.NewServer(

		retrieveStaffEndpoint,
		staffsvc.DecodeRetrieveStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveAllStaffHandler := httptransport.NewServer(

		retrieveAllStaffEndpoint,
		staffsvc.DecodeRetrieveAllStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	addRoleHandler := httptransport.NewServer(

		addRoleEndpoint,
		staffsvc.DecodeAddStaffRoleRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveRoleHandler := httptransport.NewServer(

		retrieveRoleEndpoint,
		staffsvc.DecodeRetrieveStaffRoleRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	recordBestPerformingStaffHandler := httptransport.NewServer(

		recordBestPerformingStaffEndpoint,
		staffsvc.DecodeRecordBestPerformingStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	updateBestPerformingStaffHandler := httptransport.NewServer(

		updateBestPerformingStaffEndpoint,
		staffsvc.DecodeRecordBestPerformingStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	deleteBestPerformingStaffHandler := httptransport.NewServer(

		deleteBestPerformingStaffEndpoint,
		staffsvc.DecodeRecordBestPerformingStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	recordBestPerformingStudentHandler := httptransport.NewServer(

		bestPerformingStudentEndpoint,
		staffsvc.DecodeRecordBestPerformingStudentRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	updateBestPerformingStudentHandler := httptransport.NewServer(

		updatebestPerformingStudentEndpoint,
		staffsvc.DecodeRecordBestPerformingStudentRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	deleteBestPerformingStudentHandler := httptransport.NewServer(

		deletebestPerformingStudentEndpoint,
		staffsvc.DecodeRecordBestPerformingStudentRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveBestPerformingStaffHandler := httptransport.NewServer(

		retrieveBestPerformingStaffEndpoint,
		staffsvc.DecodeRetrieveBestPerformingStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	retrieveTeacherStaffHandler := httptransport.NewServer(

		retrieveTeacherEndpoint,
		staffsvc.DecodeRetrieveStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

	retrieveStudentHandler := httptransport.NewServer(

		retrieveStudentEndpoint,
		staffsvc.DecodeRetrieveStaffRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveBestPerformingStudentHandler := httptransport.NewServer(

		retrieveBestPerformingStudentEndpoint,
		staffsvc.DecodeRetrieveBestPerformingStudentRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	rankStaffPerformanceHandler := httptransport.NewServer(

		rankStaffPerformanceEndpoint,
		staffsvc.DecodeRankStaffPerformanceRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	rankStudentPerformanceHandler := httptransport.NewServer(

		rankStudentPerformanceEndpoint,
		staffsvc.DecodeRankStudentPerformanceRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)

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

	loginUserHandler := httptransport.NewServer(
		usersvc.MakeLoginEndpoint(user),
		usersvc.DecodeLoginRequest,
		usersvc.EncodeResponse,
		options...,
	)

	//activity handlers
	//	createActivityHandler := httptransport.NewServer(
	//
	//	)

	var routes = Routes{
		Route{
			"User",
			"POST",
			"/user",
			userHandler,
		},
		Route{
			"User",
			"PATCH",
			"/user",
			updateUserHandler,
		},Route{
			"User",
			"DELETE",
			"/user",
			deleteUserHandler,
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
			"School",
			"PATCH",
			"/school",
			updateSchoolHandler,
		},
		Route{
			"School",
			"DELETE",
			"/school",
			deleteSchoolHandler,
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
			"/school",
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
		//activity routes
		Route{
			"Activity",
			"POST",
			"/activity",
			createActivityHandler,
		},

		Route{
			"Activity",
			"PATCH",
			"/activity",
			updateActivityHandler,
		},
		Route{
			"Activity",
			"DELETE",
			"/activity",
			deleteActivityHandler,
		},
		Route{
			"Activity",
			"GET",
			"/activity",
			getAllActivitiesHandler,
		},
		Route{
			"Activity",
			"GET",
			"/activity/{id}",
			getOneActivityHandler,
		},
		Route{
			"Activity Level",
			"POST",
			"/activity_level",
			createLevelHandler,
		},
		Route{
			"Activity Level",
			"GET",
			"/activity_level",
			getAllLevelsHandler,
		},
		Route{
			"Activity Level",
			"GET",
			"/activity_level/{id}",
			getOneLevelHandler,
		},
		Route{
			"Activity Performance",
			"POST",
			"/activity_performance",
			recordactivityPerformanceHandler,
		},
		Route{
			"Activity Performance",
			"GET",
			"/activity_performance",
			getAllPerformancesHandler,
		},
		Route{
			"Activity Performance",
			"GET",
			"/activity_performance/{id}",
			getOnePerformanceHandler,
		},
		//file routes
		Route{
			"File ",
			"POST",
			"/file",
			createFileHandler,
		},
		Route{
			"File ",
			"PATCH",
			"/file",
			updateFileHandler,
		},
		Route{
			"File ",
			"DELETE",
			"/file",
			deleteFileHandler,
		},
		Route{
			"File ",
			"GET",
			"/file/{id}",
			getOneFileHandler,
		},
		Route{
			"File",
			"GET",
			"/file",
			getAllFilesHandler,
		},
		Route{
			"File ",
			"POST",
			"/file_type",
			createFileTypeHandler,
		},
		Route{
			"File ",
			"GET",
			"/file_type",
			getAllFileTypesHandler,
		},
		Route{
			"File",
			"GET",
			"/file_type/{id}",
			getOneFileTypeHandler,
		},
		//infrastructure routes
		Route{
			"Infrastructure",
			"POST",
			"/infrastructure",
			createInfHandler,
		},
		Route{
			"Infrastructure",
			"PATCH",
			"/infrastructure",
			updateInfHandler,
		},
		Route{
			"Infrastructure",
			"DELETE",
			"/infrastructure",
			deleteInfHandler,
		},
		Route{
			"Infrastructure",
			"GET",
			"/infrastructure/{id}",
			getOneInfHandler,
		},
		Route{
			"Infrastructure",
			"GET",
			"/infrastructure",
			getAllInfHandler,
		},
		Route{
			"Infrastructure",
			"POST",
			"/infrastructure_type",
			createInfTypeHandler,
		},
		Route{
			"Infrastructure",
			"GET",
			"/infrastructure_type/{id}",
			getOneInfTypeHandler,
		},
		Route{
			"Infrastructure",
			"GET",
			"/infrastructure_type",
			getAllInfTypesHandler,
		},
		//message routes
		Route{
			"Message",
			"POST",
			"/message",
			createMessageHandler,
		},

		Route{
			"Message",
			"PATCH",
			"/message",
			updateMessageHandler,
		},
		Route{
			"Message",
			"DELETE",
			"/message",
			deleteMessageHandler,
		},
		Route{
			"Message",
			"GET",
			"/message/{id}",
			getOneMessageHandler,
		},
		Route{
			"Message",
			"POST",
			"/mymessage",
			getAllMessagesHandler,
		},
		//project routes
		Route{
			"Project",
			"POST",
			"/project",
			createProjectHandler,
		},
		Route{
			"Project",
			"PATCH",
			"/project",
			updateProjectHandler,
		},
		Route{
			"Project",
			"DELETE",
			"/project",
			deleteProjectHandler,
		},
		Route{
			"Project",
			"GET",
			"/project/{id}",
			getOneProjectHandler,
		},
		Route{
			"Project",
			"GET",
			"/project",
			getAllProjectHandler,
		},
		//staff routes
		Route{
			"Staff",
			"POST",
			"/staff",
			addStaffHandler,
		},
		Route{
			"Staff",
			"PATCH",
			"/staff",
			updateStaffHandler,
		},
		Route{
			"Staff",
			"DELETE",
			"/staff",
			deleteStaffHandler,
		},
		Route{
			"Staff",
			"GET",
			"/staff/{id}",
			retrieveStaffHandler,
		},
		Route{
			"Staff",
			"GET",
			"/staff",
			retrieveAllStaffHandler,
		},
		Route{
			"Staff",
			"POST",
			"/staff_role",
			addRoleHandler,
		},
		Route{
			"Staff",
			"GET",
			"/staff_role",
			retrieveRoleHandler,
		},
		Route{
			"Best Teacher",
			"POST",
			"/teacher",
			recordBestPerformingStaffHandler,
		},
		Route{
			"Best Teacher",
			"PATCH",
			"/teacher",
			updateBestPerformingStaffHandler,
		},
		Route{
			"Best Teacher",
			"DELETE",
			"/teacher",
			deleteBestPerformingStaffHandler,
		},
		Route{
			"Best Teacher",
			"GET",
			"/teacher/{id}",
			retrieveTeacherStaffHandler,
		},
		Route{
			"Best Student",
			"POST",
			"/student",
			recordBestPerformingStudentHandler,
		},
		Route{
			"Best Student",
			"PATCH",
			"/student",
			updateBestPerformingStudentHandler,
		},
		Route{
			"Best Student",
			"DELETE",
			"/student",
			deleteBestPerformingStudentHandler,
		},

		Route{
			"Best Student",
			"GET",
			"/student/{id}",
			retrieveStudentHandler,
		},
		Route{
			"Best Teacher",
			"POST",
			"/best_teacher",
			retrieveBestPerformingStaffHandler,
		},
		Route{
			"Best Student",
			"POST",
			"/best_student",
			retrieveBestPerformingStudentHandler,
		},
		Route{
			"Best Teacher",
			"POST",
			"/rank_teacher",
			rankStaffPerformanceHandler,
		},
		Route{
			"Best Student",
			"POST",
			"/rank_student",
			rankStudentPerformanceHandler,
		},
	}
	r := APINewRouter(routes)
	handler := cors.Default().Handler(r)
	//allowedOrigins := []string{"http://localhost:8080"}
	//cors.New(cors.Options{allowedOrigins})
	version1 := r.PathPrefix("/v1").Subrouter()
	//version2 := r.PathPrefix("/v2").Subrouter()
	AddRoutes(version1, routes)
	//AddRoutes(version2,routes)
	//r.Handle()
	r.Handle("/metrics", stdprometheus.Handler())
	logger.WithFields(logrus.Fields{"msg": "HTTP", "addr": ":8000"}).Info("Everything is ready, let's go !!!")
	logger.WithFields(logrus.Fields{"msg": "Serving on port", "addr":os.Getenv("PORT")}).Info("This is the port am serving from ")
	logger.WithFields(logrus.Fields{"err": http.ListenAndServe(":"+os.Getenv("PORT"), corsHandler(handler))}).Fatal("Oops! the server crashed")
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
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET,PUT,DELETE,PATCH,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			//w.WriteHeader(204)
			logrus.Debug("I got here")
			return
		}
		h.ServeHTTP(w, r)
	})
}
