package main

import (
	"log"
	"time"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/validationService/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/validationService/server"
)

//TODO: Move to a Docker container once it is online
//TODO: Move values to a config file once we have one

const (
	port          = "3000"
	listenAddress = "localhost"
)

func main() {
	startTime := time.Now()
	timestamp := startTime.Format("2006-01-02 15:04:05.000000")
	log.Println("Start: ", timestamp)

	handler := handlers.NewHandler()
	server := server.NewServer(handler, port, listenAddress)

	server.StartValidationServer(handler)

	log.Println("End: ", time.Now())
	log.Println("Uptime: ", time.Since(startTime))
}
