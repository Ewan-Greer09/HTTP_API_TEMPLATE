package main

import (
	"log"
	"time"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/config"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/server"
	"github.com/enescakir/emoji"
)

//TODO: Call the Auth service once it is online
//TODO: Move to a Database once it is online
//TODO: Move to a Docker container once it is online
//TODO: Move values to a config file once we have one
//TODO: Add a logger once we have one

const (
	port       = ":8080"
	listenAddr = "localhost"
)

func main() {
	startTime := time.Now()
	timestamp := startTime.Format("2006-01-02 15:04:05.000000")
	log.Println(emoji.StarStruck, "Start: ", timestamp)

	cfg := config.Init()

	JobBoardHandler := handlers.NewHandler(cfg)
	authHandler := auth.NewAuthHandler()
	server := server.NewServer(JobBoardHandler, authHandler, cfg.Port, cfg.Address)

	server.StartServer()

	log.Println("End: ", time.Now())
	log.Println("Uptime: ", time.Since(startTime))
}
