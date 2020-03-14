package main

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/personal/mygo/internal"
)

func main()  {
	cfg := internal.Config{}
	service := internal.NewService(zap.L(),NewHttpClient(),cfg)
	controller:= internal.NewController(zap.L(),service)
	x:= controller.RegisterRoute()
	http.ListenAndServe(":8080",x)
}

func NewHttpClient() http.Client {
	return http.Client{
		Timeout:       10*time.Second,
	}
}