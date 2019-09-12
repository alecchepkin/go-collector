package app

import (
	"github.com/go-pg/pg"
	"pg-ext"
	"sync"
)

const dbRawSchema = "collector"

var dbOnce sync.Once
var dbClient *pg.DB

func Db() *pg.DB {
	dbOnce.Do(func() {
		dbClient = pg.Connect(pg_ext.ConnOptsFromDsn(Config().DbDsn))
		if _, err := dbClient.Exec("set search_path=?", dbRawSchema); err != nil {
			Log().Panic(err)
		}
		dbClient.AddQueryHook(pg_ext.DbLogger{
			LogFunc: func(q string, p []interface{}) { Log().Debugf("query: %s", q) },
			ErrFunc: func(err error) { Log().Panic(err) },
		})
	})
	return dbClient
}

const dbServiceScheme = "banner_collector"

var dbServiceOnce sync.Once
var dbServiceClient *pg.DB

func DbService() *pg.DB {

	dbServiceOnce.Do(func() {
		dbServiceClient = pg.Connect(pg_ext.ConnOptsFromDsn(Config().DbServiceDsn))
		if _, err := dbServiceClient.Exec("set search_path=?", dbServiceScheme); err != nil {
			Log().Panic(err)
		}
		dbServiceClient.AddQueryHook(pg_ext.DbLogger{
			LogFunc: func(q string, p []interface{}) { Log().Debugf("query: %s", q) },
			ErrFunc: func(err error) { Log().Panic(err) },
		})
	})

	return dbServiceClient
}
