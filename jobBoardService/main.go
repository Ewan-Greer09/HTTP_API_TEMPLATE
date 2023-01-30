package main

import (
	"log"
	"time"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/jobBoardService/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/jobBoardService/server"
)

//TODO: Call the validation service once it is online
//TODO: Call the Auth service once it is online
//TODO: Move to a Database once it is online
//TODO: Move to a Docker container once it is online
//TODO: Move values to a config file once we have one

func main() {
	startTime := time.Now()
	timestamp := startTime.Format("2006-01-02 15:04:05.000000")
	log.Println("Start: ", timestamp)

	handler := handlers.NewHandler()

	server.StartServer(handler)

	log.Println("End: ", time.Now())
	log.Println("Uptime: ", time.Since(startTime))
}
