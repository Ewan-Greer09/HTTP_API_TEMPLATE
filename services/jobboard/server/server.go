package server

import (
	"log"
	"net/http"

	"github.com/enescakir/emoji"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/handlers"
)

type Server struct {
	Port          string
	ListenAddress string
	logger        *logger.Logger
	Handler       *handlers.Handler
	AuthHandler   *auth.AuthHandler
	db            *repository.SQLDatabase
}

func NewServer(h *handlers.Handler, auth *auth.AuthHandler, db *repository.SQLDatabase, logger *logger.Logger, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		Handler:       h,
		logger:        logger,
		ListenAddress: listenAddr,
		AuthHandler:   auth,
		db:            db,
	}
}

func (s *Server) StartServer() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.Mount("/api", s.Routes(s.db))

	s.logger.Info(emoji.Sprintf("Server started on %s:%s", s.ListenAddress, s.Port))
	log.Fatal(http.ListenAndServe(s.ListenAddress+":"+s.Port, router))
}

func (s *Server) Routes(storage *repository.SQLDatabase) http.Handler {
	r := chi.NewRouter()
	r.Post("/createlisting", s.Handler.HandleCreateListing(s.db))
	r.Get("/getlistingbyid/{id}", s.Handler.HandleGetListingByID(s.db))
	r.Post("/updatelistingbyid/{id}", s.Handler.HandleUpdateListingByID(s.db))
	r.Delete("/deletelistingbyid/{id}", s.Handler.HandleDeleteListingByID(s.db))

	r.Mount("/auth", s.AuthHandler.Routes())
	return r
}
