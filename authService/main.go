package main

import (
	"log"
	"time"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/authService/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/authService/server"
)

const (
	port       = ":9090"
	listenAddr = "localhost"
)

func main() {
	startTime := time.Now()
	timestamp := startTime.Format("2006-01-02 15:04:05.000000")
	log.Println("Start: ", timestamp)

	handler := handlers.NewHandler()
	server := server.NewServer(handler, port, listenAddr)

	server.StartServer(handler)

	log.Println("End: ", time.Now())
	log.Println("Uptime: ", time.Since(startTime))

}
