package store

import (
	dbsql "database/sql"
	l4g "github.com/alecthomas/log4go"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"salv_prj/model"
	"os"
	"strings"
	"time"
)

const (
	INDEX_TYPE_FULL_TEXT = "full_text"
	INDEX_TYPE_DEFAULT   = "default"
	MAX_DB_CONN_LIFETIME = 60
)

const (
	EXIT_CREATE_TABLE               = 100
	EXIT_DB_OPEN                    = 101
	EXIT_PING                       = 102
	EXIT_NO_DRIVER                  = 103
	EXIT_TABLE_EXISTS               = 104
	EXIT_TABLE_EXISTS_MYSQL         = 105
	EXIT_COLUMN_EXISTS              = 106
	EXIT_DOES_COLUMN_EXISTS_MYSQL   = 108
	EXIT_DOES_COLUMN_EXISTS_MISSING = 109
	EXIT_CREATE_COLUMN_MYSQL        = 111
	EXIT_CREATE_COLUMN_MISSING      = 112
	EXIT_REMOVE_COLUMN              = 113
	EXIT_RENAME_COLUMN              = 114
	EXIT_MAX_COLUMN                 = 115
	EXIT_ALTER_COLUMN               = 116
	EXIT_CREATE_INDEX_MYSQL         = 118
	EXIT_CREATE_INDEX_FULL_MYSQL    = 119
	EXIT_CREATE_INDEX_MISSING       = 120
	EXIT_REMOVE_INDEX_MYSQL         = 122
	EXIT_REMOVE_INDEX_MISSING       = 123
)

type SqlStore struct {
	master        *gorp.DbMap
	replicas      []*gorp.DbMap
	user          UserStore
	SchemaVersion string
	rrCounter     int64
}

func InitConnection() *SqlStore {
	sqlStore := &SqlStore{
		rrCounter: 0,
	}

	ConfigureApp(os.Getenv("GO_ENV"))
	//ConfigureApp("staging")
	Cache = InitCache(CACHE_URL,CACHE_PASSWORD,CACHE_DB)

	//fmt.Println("This is the env var as it is set",os.Getenv("GO_ENV"))
	/*env := os.Getenv("GO_ENV")
	if env == "staging" {
		user = "staging"
		password = "cNR6-kZhU-D68T-xa7_"
		url = "@tcp(staging-db-54f17a7d.sendyit.com:3306)/parcel?parseTime=true"
		redis_url = "redis:6379"

	} else if env == "production" {

		user = "dbuser"
		password = "ziana12345"
		url = "@tcp(5.189.153.58:3306)/parcel?parseTime=true"
		redis_url = "fe-prod-sessions.swnftp.ng.0001.euw1.cache.amazonaws.com:6379"

	}*/
	URL := "@tcp("+DB_URL+")/"+DB_DATASTORE+"?parseTime=true"
	//fmt.Println(user,password,url,redis_url)
	sqlStore.master = setupConnection("master", DB_DRIVER,
		DB_USER+":"+DB_PASSWORD+URL, DB_MAX_IDLE,
		DB_POOL, true)


	return sqlStore
}

func setupConnection(con_type string, driver string, dataSource string, maxIdle int, maxOpen int, trace bool) *gorp.DbMap {

	db, err := dbsql.Open(driver, dataSource)
	if err != nil {
		l4g.Critical("store.sql.open_conn.critical", err)
		time.Sleep(time.Second)
		os.Exit(EXIT_DB_OPEN)
	}

	l4g.Info("store.sql.pinging.info", con_type)
	err = db.Ping()
	if err != nil {
		l4g.Info("store.sql.ping.critical", err)
		time.Sleep(time.Second)
		os.Exit(EXIT_PING)
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	db.SetConnMaxLifetime(time.Duration(MAX_DB_CONN_LIFETIME) * time.Minute)

	var dbmap *gorp.DbMap
	dbmap = &gorp.DbMap{Db: db, TypeConverter:userConverter{} , Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}
	//dbmap.AddTableWithName(model.Insurer{}, "insurer").SetKeys(true, "insurer_id")
	dbmap.AddTableWithName(model.User{}, "user").SetKeys(true, "user_id")
	dbmap.AddTableWithName(model.School{}, "school").SetKeys(true, "school_id")
	//dbmap.AddTableWithName(model.Role{}, "insurer_user_role").SetKeys(true, "insurer_user_role_id")

	return dbmap
}

func (ss *SqlStore) GetMaster() *gorp.DbMap {
	return ss.master
}

func (ss *SqlStore) TotalMasterDbConnections() int {
	return ss.GetMaster().Db.Stats().OpenConnections
}

func (ss *SqlStore) Close() {
	l4g.Info("store.sql.closing.info")
	ss.master.Db.Close()
	//for _, replica := range ss.replicas {
	//	replica.Db.Close()
	//}
}

func IsUniqueConstraintError(err string, indexName []string) bool {
	unique := strings.Contains(err, "unique constraint") || strings.Contains(err, "Duplicate entry")
	field := false
	for _, contain := range indexName {
		if strings.Contains(err, contain) {
			field = true
			break
		}
	}

	return unique && field
}

//func (ss *SqlStore) Close(){
//	ss.master.Db.Close()
//}


