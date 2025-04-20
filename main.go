package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tdidf/internal/config"
	"tdidf/internal/handler"
	"tdidf/internal/service"
)

func main() {
	route := gin.Default()
	conf := config.FromFile("config.yaml")
	serv := service.New(conf)
	hndlr := handler.NewHandler(route, serv)
	hndlr.Init()

	err := route.Run(conf.ListenPort)
	if err != nil {
		log.Println(err)
		return
	}
}
