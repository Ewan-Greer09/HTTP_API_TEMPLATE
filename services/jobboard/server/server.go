package server

import (
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
	JobHandler    *handlers.Handler
	AuthHandler   *auth.AuthHandler
	db            *repository.GormDatabase
}

func NewServer(h *handlers.Handler, auth *auth.AuthHandler, db *repository.GormDatabase, logger *logger.Logger, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		JobHandler:    h,
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

	router.Mount("/api", s.AuthHandler.VerifyJWT(s.Routes()))

	s.logger.Info(emoji.Sprintf("Server started on %s:%s", s.ListenAddress, s.Port))
	s.logger.Panic(http.ListenAndServe(s.ListenAddress+":"+s.Port, router))
}

func (s *Server) Routes() http.HandlerFunc {
	r := chi.NewRouter()

	r.Post("/listing", s.JobHandler.HandleCreateListing(s.db))
	r.Get("/listing/{id}", s.JobHandler.HandleGetListingByID(s.db))
	r.Post("/listing/{id}", s.JobHandler.UpdateJobListing(s.db))
	r.Delete("/listing/{id}", s.JobHandler.HandleDeleteListingByID(s.db))
	r.Get("/listing", s.JobHandler.HandleAllListings(s.db))

	r.Mount("/auth", s.AuthHandler.Routes())
	return r.ServeHTTP
}
