package store

import (

	"github.com/spf13/viper"
	l4g "github.com/alecthomas/log4go"
)

var(
	DB_USER = "root"
	DB_PASSWORD = ""
	DB_URL = "localhost:3306"
	DB_DATASTORE = "parcel"
	DB_DRIVER = "mysql"
	DB_POOL = 10
	DB_MAX_IDLE = 3
	CACHE_USER = ""
	CACHE_PASSWORD = ""
	CACHE_URL = "localhost:6379"
	CACHE_DB = 0

)

func ConfigureApp(env string) bool {

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AddConfigPath("/config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()

	if err != nil {
		l4g.Critical("No configuration file loaded - using defaults",err.Error())
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
	}else if env == "production" {


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

	}

	l4g.Info("Am done configuring your app, time to roll!!")
	return true

}
