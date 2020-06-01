package cache

import (
	"log"
	"testing"

	"github.com/gmbanaag/wallet2020/internal/app/config"
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"

	. "github.com/smartystreets/goconvey/convey"
)

var cacheService *Service

func init() {
	if err := godotenv.Load("../../../../.envtest"); err != nil {
		log.Fatal(err)
	}
	cfg := &config.Config{}
	if err := envconfig.Init(&cfg); err != nil {
		log.Fatal(err)
	}

	cacheService = &Service{}
	cacheService.InitClient(cfg.RedisAddr, cfg.RedisKey)
}
func TestCache(t *testing.T) {
	Convey("Testing cache connection", t, func() {
		Convey("get key", func() {
			err := cacheService.Set("test12345", "test1", 0)
			if err != nil {
				log.Println(err)
			}
			value, _ := cacheService.Get("test12345")
			So(value, ShouldEqual, "test1")
		})

		Convey("set key", func() {
			_ = cacheService.Set("test12345", "test2", 0)
			value, _ := cacheService.Get("test12345")
			So(value, ShouldEqual, "test2")
		})

		Convey("delete key", func() {
			delval, _ := cacheService.Delete("test12345")
			value, _ := cacheService.Get("test12345")
			So(delval, ShouldEqual, 1)
			So(value, ShouldBeEmpty)
		})
	})
}
