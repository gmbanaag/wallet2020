package model

import (
	"log"

	"github.com/gmbanaag/wallet2020/internal/app/config"
	"github.com/gmbanaag/wallet2020/internal/app/model/cache"
	uuid "github.com/satori/go.uuid"

	//"github.com/go-chi/chi/middleware"
	//"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

var cfg *config.Config

func initTest() {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err)
	}
	cfg := &config.Config{}
	if err := envconfig.Init(&cfg); err != nil {
		log.Fatal(err)
	}

	InitializeDB(cfg.MysqlConnW)

	cacheService := &cache.Service{}
	cacheService.InitClient(cfg.RedisAddr, cfg.RedisKey)
	SetCacheClient(cacheService)

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Wallet{}, &Transaction{})
}

func generateUUID() string {
	id := uuid.NewV4()

	return id.String()
}
