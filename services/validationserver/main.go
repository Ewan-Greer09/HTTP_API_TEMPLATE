package main

import (
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/validationserver/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/validationserver/server"
	"github.com/enescakir/emoji"
)

//TODO: Move to a Docker container once it is online
//TODO: Move values to a config file once we have one

const (
	port          = ":3000"
	listenAddress = "localhost"
)

func main() {
	logger := logger.NewLogger()
	logger.SetOutput("validationserver.log")

	logger.Info(emoji.Sprint("Starting Validation Server :rocket:"))

	handler := handlers.NewHandler()
	server := server.NewServer(handler, logger, port, listenAddress)

	logger.Info(emoji.Sprintf("Validation Server started on %s:%s :rocket:", listenAddress, port))
	server.StartValidationServer()

}
