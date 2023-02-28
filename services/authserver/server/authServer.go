package server

import (
	"log"
	"net/http"
	"time"


	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/authserver/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Port          string
	ListenAddress string
	Handler       *handlers.Handler
}

func NewServer(h *handlers.Handler, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		Handler:       h,
		ListenAddress: listenAddr,
	}
}

func (s *Server) StartServer(h *handlers.Handler) {
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Starting HTTP API Auth Service...")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Get("/ping", h.HandlePing)
		r.Post("/authenticate", h.HandleAuthenticate)
	})

	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Listening and serving on port "+s.Port)
	log.Fatal(http.ListenAndServe(s.Port, router))
}