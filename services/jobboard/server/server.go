package server

import (
	"net/http"

	"github.com/enescakir/emoji"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/storage"
)

type Server struct {
	Port          string
	ListenAddress string
	logger        *logger.Logger
	JobHandler    *handlers.Handler
	AuthHandler   *auth.AuthHandler
}

func NewServer(h *handlers.Handler, auth *auth.AuthHandler, logger *logger.Logger, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		JobHandler:    h,
		logger:        logger,
		ListenAddress: listenAddr,
		AuthHandler:   auth,
	}
}

func (s *Server) StartServer() {
	s.logger.Info(emoji.Sprint("Starting server..."))
	s.logger.Info(emoji.Sprint("Populating storage..."))
	storage := storage.PopulateStorage()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.Mount("/api", s.Routes(storage))

	s.logger.Info(emoji.Sprintf("Server started on %s:%s", s.ListenAddress, s.Port))
	s.logger.Fatal(http.ListenAndServe(s.ListenAddress+":"+s.Port, router))
}

func (s *Server) Routes(storage map[string]types.JobListing) http.Handler {
	r := chi.NewRouter()
	r.Post("/listing", s.JobHandler.HandleCreateListing(storage))
	r.Get("/listing/{id}", s.JobHandler.HandleGetListingByID(storage))
	r.Post("/listing/{id}", s.JobHandler.HandleUpdateListingByID(storage))
	r.Delete("/listing/{id}", s.JobHandler.HandleDeleteListingByID(storage))

	r.Mount("/auth", s.AuthHandler.Routes())
	return r
}
