package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/henglory/Demo_Golang_v0.0.1/aclient"
	"github.com/henglory/Demo_Golang_v0.0.1/bclient"
	"github.com/henglory/Demo_Golang_v0.0.1/config"
	"github.com/henglory/Demo_Golang_v0.0.1/logger"
	"github.com/henglory/Demo_Golang_v0.0.1/server"
	"github.com/henglory/Demo_Golang_v0.0.1/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	acli := aclient.New(config.ActionAUrl, config.ATimeout, logger.WriteLog)
	bcli := bclient.New(config.ActionBUrl, config.BTimeout, logger.WriteLog)

	s := service.Service{
		ARequestFn: acli.Request,
		BRequestFn: bcli.Request,
	}

	log.Println("Start server")
	serv := server.NewServer(s)
	serv.Start()
	defer serv.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")
}
