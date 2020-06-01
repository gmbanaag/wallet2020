package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gmbanaag/wallet2020/internal/app/config"
	"github.com/gmbanaag/wallet2020/internal/app/handler"
	"github.com/gmbanaag/wallet2020/internal/app/logger"
	md "github.com/gmbanaag/wallet2020/internal/app/middleware"
	"github.com/gmbanaag/wallet2020/internal/app/model"
	"github.com/gmbanaag/wallet2020/internal/app/model/cache"
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var dbServiceR *sql.DB
var dbServiceW *sql.DB
var router *chi.Mux
var cfg *config.Config

func main() {
	LoadConfig()
	SetLogging()

	router = chi.NewRouter()
	LoadMiddleWares()
	InitializeModels()
	LoadHandlers()
	StartWeb()
}

//SetLogging set logging options
func SetLogging() {
	log.SetFlags(log.Ldate | log.Ltime)
	logger.SetLogLevel(cfg.LogLevel)
}

//LoadMiddleWares sets middlewares
func LoadMiddleWares() {

	mwAuth := md.Auth{Config: cfg}
	router.Use(mwAuth.Authenticate)

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	logger.LogInfo("Loading middlewares are done")
}

//LoadConfig parse service configuration
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	cfg = &config.Config{}
	if err := envconfig.Init(&cfg); err != nil {
		log.Fatal(err)
	}

	logger.LogInfo("Loading env are done")
}

//InitializeModels initialize db and cache connections
func InitializeModels() {
	model.InitializeDB(cfg.MysqlConnW)

	cacheService := &cache.Service{}
	cacheService.InitClient(cfg.RedisAddr, cfg.RedisKey)
	model.SetCacheClient(cacheService)

	model.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Wallet{}, &model.Transaction{})
	logger.LogInfo("Initializing models are done")
}

//LoadHandlers initialize controllers`
func LoadHandlers() {
	hd := handler.NewHandler(cfg)

	handler.LoadRouters(router, hd)

	logger.LogInfo("Loading routes and controllers are done")
}

//StartServer starts the service
func StartServer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r) // dispatch the request
	})
}

//StartWeb start http listener
func StartWeb() {
	listenAddress := fmt.Sprintf("0.0.0.0:%s", cfg.Port)

	logger.LogInfo(fmt.Sprintf("Server listening to: %s", listenAddress))

	err := http.ListenAndServe(listenAddress, StartServer())
	catch(err, true)
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
