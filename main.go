package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/henglory/Demo_Golang_v0.0.1/server"
	"github.com/henglory/Demo_Golang_v0.0.1/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s := service.Service{}

	log.Println("Start server")
	serv := server.NewServer(s)
	serv.Start()
	defer serv.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")
}
