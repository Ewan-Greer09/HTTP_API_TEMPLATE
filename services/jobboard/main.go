package main

import (
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/config"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/server"
	"github.com/enescakir/emoji"
)

//TODO: Move to a Database once it is online
//TODO: Move to a Docker container once it is online

func main() {
	logger := logger.NewLogger()
	logger.SetOutput("jobboard.log")

	logger.Info(emoji.Sprint("Starting Job Board Service :rocket:"))

	cfg := config.Init()

	logger.Info(emoji.Sprintf("Creating local database :rocket:"))
	db, err := repository.NewDatabase(logger)
	if err != nil {
		logger.Errorf("Failed to create database: %v", err)
		return
	}
	logger.Info(emoji.Sprintf("Database created :rocket:"))

	JobBoardHandler := handlers.NewHandler(cfg)
	authHandler := auth.NewAuthHandler()
	server := server.NewServer(JobBoardHandler, authHandler, db, logger, cfg.Port, cfg.Address)

	logger.Info(emoji.Sprintf("Job Board Service started on %s:%s :rocket:", cfg.Address, cfg.Port))
	server.StartServer()
}
