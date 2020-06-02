package model

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/gmbanaag/wallet2020/internal/app/logger"
	"github.com/gmbanaag/wallet2020/internal/app/model/cache"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB handler
var DB *gorm.DB

//CacheClient hold the client connection
var CacheClient *cache.Service

var cacheOnce sync.Once
var dbOnce sync.Once

const (
	//LimitDefault states the limit if theres nothing stated in the request
	LimitDefault = 100
	//LimitMax row count limit should not exceed this value
	LimitMax = 1000
)

//InitializeDB will warm-up DB connection/s
func InitializeDB(connectionString string) {
	setDB := func() {
		db, err := gorm.Open("mysql", connectionString)
		if err != nil {
			panic(err)
		}

		db.DB()
		err = db.DB().Ping()
		if err != nil {
			panic(err)
		}

		db.DB().SetMaxIdleConns(0)
		db.DB().SetMaxOpenConns(151)
		DB = db
	}

	dbOnce.Do(setDB)
}

//SetCacheClient warms-up cache connection
func SetCacheClient(client *cache.Service) {
	setCache := func() {
		CacheClient = client
	}

	cacheOnce.Do(setCache)
}

func catch(err error, fatal bool) {
	_, fn, line, _ := runtime.Caller(1)
	message := fmt.Sprintf("%s:%d %v", fn, line, err)

	if fatal == true && err != nil {
		logger.LogFatal(message)
		panic(err)
	}

	if err != nil {
		logger.LogError(message)
	}
}
