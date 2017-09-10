package store

import (
	l4g "github.com/alecthomas/log4go"
	"github.com/spf13/viper"
)

//var (
//	DB_USER        = "c25ccz1z5g278kqu"
//	DB_PASSWORD    = "ixjg7e9rw4c3tmda"
//	DB_URL         = "p2d0untihotgr5f6.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306"
//	DB_DATASTORE   = "ev7x2lf58ardm8lg"
//	DB_DRIVER      = "mysql"
//	DB_POOL        = 10
//	DB_MAX_IDLE    = 3
//	CACHE_USER     = "redistogo"
//	CACHE_PASSWORD = "a91baadf53ef530bccfeaa71ebab2aaf"
//	CACHE_URL      = "angelfish.redistogo.com:11928"
//	CACHE_DB       = 0
//)
var (
	DB_USER        = ""
	DB_PASSWORD    = ""
	DB_URL         = ""
	DB_DATASTORE   = ""
	DB_DRIVER      = ""
	DB_POOL        = 10
	DB_MAX_IDLE    = 3
	CACHE_USER     = ""
	CACHE_PASSWORD = ""
	CACHE_URL      = ""
	CACHE_DB       = 0
)
//
//"local": {
//"ConnectionUrl":"p2d0untihotgr5f6.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306",
//"User":"c25ccz1z5g278kqu",
//"Password":"ixjg7e9rw4c3tmda",
//"Database":"ev7x2lf58ardm8lg",
//"Driver":"mysql",
//"ConnectionPool":10,
//"MaxIdleConnections":3
//},

func ConfigureApp(env string) bool {

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AddConfigPath("/config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()

	if err != nil {
		l4g.Critical("No configuration file loaded - using defaults", err.Error())
		return false
	}
	if env == "staging" {
		//load DB configs
		DB_USER = viper.GetString("Databases.staging.User")
		DB_PASSWORD = viper.GetString("Databases.staging.Password")
		DB_URL = viper.GetString("Databases.staging.ConnectionUrl")
		DB_DATASTORE = viper.GetString("Databases.staging.Database")
		DB_DRIVER = viper.GetString("Databases.staging.Driver")
		DB_POOL = viper.GetInt("Databases.staging.ConnectionPool")
		DB_MAX_IDLE = viper.GetInt("Databases.staging.MaxIdleConnections")
		//load cache configs
		CACHE_USER = viper.GetString("Cache.staging.User")
		CACHE_PASSWORD = viper.GetString("Cache.staging.Password")
		CACHE_URL = viper.GetString("Cache.staging.ConnectionUrl")
		CACHE_DB = viper.GetInt("Cache.staging.Database")

		//configure server
	} else if env == "production" {

		//load DB configs
		DB_USER = viper.GetString("Databases.production.User")
		DB_PASSWORD = viper.GetString("Databases.production.Password")
		DB_URL = viper.GetString("Databases.production.ConnectionUrl")
		DB_DATASTORE = viper.GetString("Databases.production.Database")
		DB_DRIVER = viper.GetString("Databases.production.Driver")
		DB_POOL = viper.GetInt("Databases.production.ConnectionPool")
		DB_MAX_IDLE = viper.GetInt("Databases.production.MaxIdleConnections")
		//load cache configs
		CACHE_USER = viper.GetString("Cache.production.User")
		CACHE_PASSWORD = viper.GetString("Cache.production.Password")
		CACHE_URL = viper.GetString("Cache.production.ConnectionUrl")
		CACHE_DB = viper.GetInt("Cache.production.Database")

		//configure server

	} else if env == "heroku" {

	//load DB configs
	DB_USER = viper.GetString("Databases.heroku.User")
	DB_PASSWORD = viper.GetString("Databases.heroku.Password")
	DB_URL = viper.GetString("Databases.heroku.ConnectionUrl")
	DB_DATASTORE = viper.GetString("Databases.heroku.Database")
	DB_DRIVER = viper.GetString("Databases.heroku.Driver")
	DB_POOL = viper.GetInt("Databases.heroku.ConnectionPool")
	DB_MAX_IDLE = viper.GetInt("Databases.heroku.MaxIdleConnections")
	//load cache configs
	CACHE_USER = viper.GetString("Cache.heroku.User")
	CACHE_PASSWORD = viper.GetString("Cache.heroku.Password")
	CACHE_URL = viper.GetString("Cache.heroku.ConnectionUrl")
	CACHE_DB = viper.GetInt("Cache.heroku.Database")

	//configure server

	} else {

	//load DB configs
	DB_USER = viper.GetString("Databases.local.User")
	DB_PASSWORD = viper.GetString("Databases.local.Password")
	DB_URL = viper.GetString("Databases.local.ConnectionUrl")
	DB_DATASTORE = viper.GetString("Databases.local.Database")
	DB_DRIVER = viper.GetString("Databases.local.Driver")
	DB_POOL = viper.GetInt("Databases.local.ConnectionPool")
	DB_MAX_IDLE = viper.GetInt("Databases.local.MaxIdleConnections")
	//load cache configs
	CACHE_USER = viper.GetString("Cache.local.User")
	CACHE_PASSWORD = viper.GetString("Cache.local.Password")
	CACHE_URL = viper.GetString("Cache.local.ConnectionUrl")
	CACHE_DB = viper.GetInt("Cache.local.Database")

	//configure server

	}

	l4g.Info("Am done configuring your app, time to roll!!")
	return true

}
