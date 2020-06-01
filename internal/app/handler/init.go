package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/gmbanaag/wallet2020/internal/app/config"
	"github.com/gmbanaag/wallet2020/internal/app/logger"
)

var reqID interface{}

//Handler object
type Handler struct {
	cfg *config.Config
}

//NewHandler instantiates the Handler object
func NewHandler(cfg *config.Config) *Handler {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Handler{
		cfg: cfg,
	}
}

//Response object
type Response struct {
	Code    int `json:"-"`
	Payload interface{}
}

//Render placeholder
func (resp Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
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
