package handler

import (
	"log"

	"github.com/gmbanaag/wallet2020/internal/app/config"
	"github.com/gmbanaag/wallet2020/internal/app/mocks"
	"github.com/gmbanaag/wallet2020/internal/app/model"
	"github.com/gmbanaag/wallet2020/internal/app/model/cache"

	"github.com/go-chi/chi"
	//"github.com/go-chi/chi/middleware"
	//"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)


var cfg *config.Config
var testrouter *chi.Mux
var handler *Handler

func init() {
	if err := godotenv.Load("../../../.envtest"); err != nil {
		log.Fatal(err)
	}
	cfg := &config.Config{}
	if err := envconfig.Init(&cfg); err != nil {
		log.Fatal(err)
	}

	model.InitializeDB(cfg.MysqlConnW)

	cacheService := &cache.Service{}
	cacheService.InitClient(cfg.RedisAddr, cfg.RedisKey)
	model.SetCacheClient(cacheService)

	model.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Wallet{}, &model.Transaction{})

	testrouter = chi.NewRouter()
	handler = NewHandler(cfg)
	
	mwAuth := mocks.MockAuth{Config: cfg}
    testrouter.Use(mwAuth.Authenticate)

	LoadRouters(testrouter, handler)
}