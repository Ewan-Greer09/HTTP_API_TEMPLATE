package server

import (
	"net/http"
	"time"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/validationserver/handlers"
	"github.com/enescakir/emoji"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Port          string
	ListenAddress string
	Handler       *handlers.Handler
	Logger        *logger.Logger
}

func NewServer(h *handlers.Handler, logger *logger.Logger, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		Handler:       h,
		ListenAddress: listenAddr,
		Logger:        logger,
	}
}

func (s *Server) StartValidationServer() {
	s.Logger.Info(emoji.Sprint("Starting Validation Server :rocket:"))

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Post("/validate", s.Handler.HandleValidate)
	})

	s.Logger.Info(emoji.Sprintf("Validation Server started on %s:%s :rocket:", s.ListenAddress, s.Port))
	s.Logger.Panic(http.ListenAndServe(s.ListenAddress+s.Port, router))
}
