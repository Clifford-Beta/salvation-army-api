package main

import (
	//"context"
	"crypto/subtle"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	//customersvc "salv_prj/customerservice"
	//ordersvc "salv_prj/orderservice"
	//partnersvc "salv_prj/partnerservice"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	actsvc "salv_prj/activity"
	"salv_prj/auth"
	catsvc "salv_prj/category"
	filesvc "salv_prj/file"
	infsvc "salv_prj/infrastructure"
	msgsvc "salv_prj/message"
	projectsvc "salv_prj/project"
	schoolsvc "salv_prj/school"
	staffsvc "salv_prj/staff"
	"salv_prj/store"
	usersvc "salv_prj/user"
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
		userCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(userCreateEndpoint)
	}
	var getOneUserEndpoint endpoint.Endpoint
	{
		getOneUserEndpoint = usersvc.MakeGetOneEndpoint(user)
		getOneUserEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneUserEndpoint)
	}

	var loginUserEndpoint endpoint.Endpoint
	{
		loginUserEndpoint = usersvc.MakeLoginEndpoint(user)
		loginUserEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(loginUserEndpoint)
	}
	var getAllUsersEndpoint endpoint.Endpoint
	{
		getAllUsersEndpoint = usersvc.MakeGetAllEndpoint(user)
		getAllUsersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllUsersEndpoint)
	}
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

	//category endpoint

	var categoryCreateEndpoint endpoint.Endpoint
	{
		categoryCreateEndpoint = catsvc.MakeCreateEndpoint(category)
		categoryCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(categoryCreateEndpoint)
	}
	var getOneCategoryEndpoint endpoint.Endpoint
	{
		getOneCategoryEndpoint = catsvc.MakeGetOneEndpoint(category)
		getOneCategoryEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneCategoryEndpoint)
	}
	var getAllCategoriesEndpoint endpoint.Endpoint
	{
		getAllCategoriesEndpoint = catsvc.MakeGetAllEndpoint(category)
		getAllCategoriesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllCategoriesEndpoint)
	}
	var tierCreateEndpoint endpoint.Endpoint
	{
		tierCreateEndpoint = catsvc.MakeCreateTierEndpoint(category)
		tierCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(tierCreateEndpoint)
	}
	var getOneTierEndpoint endpoint.Endpoint
	{
		getOneTierEndpoint = catsvc.MakeGetOneTierEndpoint(category)
		getOneTierEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneTierEndpoint)
	}
	var getAllTiersEndpoint endpoint.Endpoint
	{
		getAllTiersEndpoint = catsvc.MakeGetAllTiersEndpoint(category)
		getAllTiersEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllTiersEndpoint)
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
		schoolCreateEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(schoolCreateEndpoint)
	}
	var getOneSchoolEndpoint endpoint.Endpoint
	{
		getOneSchoolEndpoint = schoolsvc.MakeGetOneEndpoint(school)
		getOneSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneSchoolEndpoint)
	}
	var getAllSchoolsEndpoint endpoint.Endpoint
	{
		getAllSchoolsEndpoint = schoolsvc.MakeGetAllEndpoint(school)
		getAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllSchoolsEndpoint)
	}

	var getBestSchoolEndpoint endpoint.Endpoint
	{
		getBestSchoolEndpoint = schoolsvc.MakeRetrieveBestPerfomingSchoolEndpoint(school)
		getBestSchoolEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getBestSchoolEndpoint)
	}

	var recordSchoolPerformanceEndpoint endpoint.Endpoint
	{
		recordSchoolPerformanceEndpoint = schoolsvc.MakeRecordPerformanceEndpoint(school)
		recordSchoolPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(recordSchoolPerformanceEndpoint)
	}
	var rankAllSchoolsEndpoint endpoint.Endpoint
	{
		rankAllSchoolsEndpoint = schoolsvc.MakeRankAllSchoolsEndpoint(school)
		rankAllSchoolsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(rankAllSchoolsEndpoint)
	}

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

	//activity endpoint
	var createActivityEndpoint endpoint.Endpoint
	{
		createActivityEndpoint = actsvc.MakeCreateActivityEndpoint(activity)
		createActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createActivityEndpoint)
	}
	var createLevelEndpoint endpoint.Endpoint
	{
		createLevelEndpoint = actsvc.MakeCreateLevelEndpoint(activity)
		createLevelEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createLevelEndpoint)
	}
	var recordPerformanceEndpoint endpoint.Endpoint
	{
		recordPerformanceEndpoint = actsvc.MakeRecordPerformanceEndpoint(activity)
		recordPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(recordPerformanceEndpoint)
	}
	var getOneActivityEndpoint endpoint.Endpoint
	{
		getOneActivityEndpoint = actsvc.MakeGetOneActivityEndpoint(activity)
		getOneActivityEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneActivityEndpoint)
	}

	var getOneLevelEndpoint endpoint.Endpoint
	{
		getOneLevelEndpoint = actsvc.MakeGetOneLevelEndpoint(activity)
		getOneLevelEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneLevelEndpoint)
	}
	var getOnePerformanceEndpoint endpoint.Endpoint
	{
		getOnePerformanceEndpoint = actsvc.MakeGetOnePerformanceEndpoint(activity)
		getOnePerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOnePerformanceEndpoint)
	}
	var getAllActivitiesEndpoint endpoint.Endpoint
	{
		getAllActivitiesEndpoint = actsvc.MakeGetAllActivitiesEndpoint(activity)
		getAllActivitiesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllActivitiesEndpoint)
	}
	var getAllLevelsEndpoint endpoint.Endpoint
	{
		getAllLevelsEndpoint = actsvc.MakeGetAllLevelsEndpoint(activity)
		getAllLevelsEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllLevelsEndpoint)
	}
	var getAllPerformancesEndpoint endpoint.Endpoint
	{
		getAllPerformancesEndpoint = actsvc.MakeGetAllPerformancesEndpoint(activity)
		getAllPerformancesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllPerformancesEndpoint)
	}

	//activity handlers
	createActivityHandler := httptransport.NewServer(
		createActivityEndpoint,
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
		createFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createFileEndpoint)
	}

	var createFileTypeEndpoint endpoint.Endpoint
	{
		createFileTypeEndpoint = filesvc.MakeCreateTypeEndpoint(file)
		createFileTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createFileTypeEndpoint)
	}

	var getOneFileEndpoint endpoint.Endpoint
	{
		getOneFileEndpoint = filesvc.MakeGetOneEndpoint(file)
		getOneFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneFileEndpoint)
	}

	var getOneFileTypeEndpoint endpoint.Endpoint
	{
		getOneFileTypeEndpoint = filesvc.MakeGetOneTypeEndpoint(file)
		getOneFileTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneFileTypeEndpoint)
	}

	var getAllFilesEndpoint endpoint.Endpoint
	{
		getAllFilesEndpoint = filesvc.MakeGetAllEndpoint(file)
		getAllFilesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllFilesEndpoint)
	}
	var getAllFileTypesEndpoint endpoint.Endpoint
	{
		getAllFileTypesEndpoint = filesvc.MakeGetAllTypesEndpoint(file)
		getAllFileTypesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllFileTypesEndpoint)
	}

	//file handlers

	createFileHandler := httptransport.NewServer(
		createFileEndpoint,
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
		createInfEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createInfEndpoint)
	}

	var createInfTypeEndpoint endpoint.Endpoint
	{
		createInfTypeEndpoint = infsvc.MakeCreateTypeEndpoint(infs)
		createInfTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createInfTypeEndpoint)
	}

	var getOneInfEndpoint endpoint.Endpoint
	{
		getOneFileEndpoint = infsvc.MakeGetOneEndpoint(infs)
		getOneFileEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneInfEndpoint)
	}

	var getOneInfTypeEndpoint endpoint.Endpoint
	{
		getOneFileTypeEndpoint = infsvc.MakeGetOneTypeEndpoint(infs)
		getOneFileTypeEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneInfTypeEndpoint)
	}

	var getAllInfEndpoint endpoint.Endpoint
	{
		getAllFilesEndpoint = infsvc.MakeGetAllEndpoint(infs)
		getAllFilesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllInfEndpoint)
	}
	var getAllInfTypesEndpoint endpoint.Endpoint
	{
		getAllFileTypesEndpoint = infsvc.MakeGetAllTypesEndpoint(infs)
		getAllFileTypesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllInfTypesEndpoint)
	}

	//infrastructure handlers
	createInfHandler := httptransport.NewServer(
		createInfEndpoint,
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
		createMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createMessageEndpoint)
	}
	var getOneMessageEndpoint endpoint.Endpoint
	{
		getOneMessageEndpoint = msgsvc.MakeGetOneEndpoint(message)
		getOneMessageEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneMessageEndpoint)
	}
	var getAllMessagesEndpoint endpoint.Endpoint
	{
		getAllMessagesEndpoint = msgsvc.MakeGetAllEndpoint(message)
		getAllMessagesEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllMessagesEndpoint)
	}

	//message handlers
	createMessageHandler := httptransport.NewServer(
		createMessageEndpoint,
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
		createProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(createProjectEndpoint)
	}
	var getOneProjectEndpoint endpoint.Endpoint
	{
		getOneProjectEndpoint = projectsvc.MakeGetOneEndpoint(project)
		getOneProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getOneProjectEndpoint)
	}
	var getAllProjectEndpoint endpoint.Endpoint
	{
		getAllProjectEndpoint = projectsvc.MakeGetAllEndpoint(project)
		getAllProjectEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(getAllProjectEndpoint)
	}

	//project handler

	createProjectHandler := httptransport.NewServer(
		createProjectEndpoint,
		projectsvc.DecodeCreateRequest,
		projectsvc.EncodeResponse,
		jwtOptions...,
	)
	getOneProjectHandler := httptransport.NewServer(
		getOneMessageEndpoint,
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
		addStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(addStaffEndpoint)
	}
	var retrieveStaffEndpoint endpoint.Endpoint
	{
		retrieveStaffEndpoint = staffsvc.MakeRetrieveStaffEndpoint(staff)
		retrieveStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(retrieveStaffEndpoint)
	}
	var retrieveAllStaffEndpoint endpoint.Endpoint
	{
		retrieveAllStaffEndpoint = staffsvc.MakeRetrieveAllStaffEndpoint(staff)
		retrieveAllStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(retrieveAllStaffEndpoint)
	}
	var addRoleEndpoint endpoint.Endpoint
	{
		addRoleEndpoint = staffsvc.MakeAddStaffRoleEndpoint(staff)
		addRoleEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(addRoleEndpoint)
	}
	var retrieveRoleEndpoint endpoint.Endpoint
	{
		retrieveRoleEndpoint = staffsvc.MakeRetrieveStaffRoleEndpoint(staff)
		retrieveRoleEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(retrieveRoleEndpoint)
	}
	var recordBestPerformingStaffEndpoint endpoint.Endpoint
	{
		recordBestPerformingStaffEndpoint = staffsvc.MakeRecordBestPerformingStaffEndpoint(staff)
		recordBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(recordBestPerformingStaffEndpoint)
	}

	var bestPerformingStudentEndpoint endpoint.Endpoint
	{
		bestPerformingStudentEndpoint = staffsvc.MakeRecordBestPerformingStudentEndpoint(staff)
		bestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(bestPerformingStudentEndpoint)
	}
	var retrieveBestPerformingStaffEndpoint endpoint.Endpoint
	{
		retrieveBestPerformingStaffEndpoint = staffsvc.MakeRetrieveBestPerformingStaffEndpoint(staff)
		retrieveBestPerformingStaffEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(retrieveBestPerformingStaffEndpoint)
	}
	var retrieveBestPerformingStudentEndpoint endpoint.Endpoint
	{
		retrieveBestPerformingStudentEndpoint = staffsvc.MakeRetrieveBestPerformingStudentEndpoint(staff)
		retrieveBestPerformingStudentEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(retrieveBestPerformingStudentEndpoint)
	}
	var rankStaffPerformanceEndpoint endpoint.Endpoint
	{
		rankStaffPerformanceEndpoint = staffsvc.MakeRankStaffPerformanceEndpoint(staff)
		rankStaffPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(rankStaffPerformanceEndpoint)
	}
	var rankStudentPerformanceEndpoint endpoint.Endpoint
	{
		rankStudentPerformanceEndpoint = staffsvc.MakeRankStudentPerformanceEndpoint(staff)
		rankStudentPerformanceEndpoint = jwt.NewParser(keys, stdjwt.SigningMethodHS256, &auth.CustomClaims{})(rankStudentPerformanceEndpoint)
	}

	//staff handlers

	addStaffHandler := httptransport.NewServer(
		addStaffEndpoint,
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
	recordBestPerformingStudentHandler := httptransport.NewServer(
		bestPerformingStudentEndpoint,
		staffsvc.DecodeRetrieveBestPerformingStudentRequest,
		staffsvc.EncodeResponse,
		jwtOptions...,
	)
	retrieveBestPerformingStaffHandler := httptransport.NewServer(
		retrieveBestPerformingStaffEndpoint,
		staffsvc.DecodeRetrieveBestPerformingStaffRequest,
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
		rankAllSchoolsEndpoint,
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
		//activity routes
		Route{
			"Activity",
			"POST",
			"/activity",
			createActivityHandler,
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
			"GET",
			"/message/{id}",
			getOneMessageHandler,
		},
		Route{
			"Message",
			"GET",
			"/message",
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
			"Best Student",
			"POST",
			"/student",
			recordBestPerformingStudentHandler,
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
	version1 := r.PathPrefix("/v1").Subrouter()
	//version2 := r.PathPrefix("/v2").Subrouter()
	AddRoutes(version1, routes)
	//AddRoutes(version2,routes)
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
		h.ServeHTTP(w, r)
	})
}
