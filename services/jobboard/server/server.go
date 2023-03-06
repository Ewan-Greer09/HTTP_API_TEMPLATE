package server

import (
	"log"
	"net/http"
	"time"

	"github.com/enescakir/emoji"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/storage"
)

type Server struct {
	Port          string
	ListenAddress string
	Handler       *handlers.Handler
	AuthHandler   *auth.AuthHandler
}

func NewServer(h *handlers.Handler, auth *auth.AuthHandler, port, listenAddr string) *Server {
	return &Server{
		Port:          port,
		Handler:       h,
		ListenAddress: listenAddr,
		AuthHandler:   auth,
	}
}

func (s *Server) StartServer() {
	log.Println(emoji.Airplane, time.Now().Format("2006-01-02 15:04:05.000000"), "Starting HTTP API Template Service...")
	log.Println(emoji.Fire, "Populating storage...")
	storage := storage.PopulateStorage()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.Mount("/api", s.Routes(storage))

	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Listening and serving on port :"+s.Port, emoji.Headphone)
	log.Fatal(http.ListenAndServe(s.ListenAddress+":"+s.Port, router))
}

func (s *Server) Routes(storage map[string]types.JobListing) http.Handler {
	r := chi.NewRouter()
	r.Post("/createlisting", s.Handler.HandleCreateListing(storage))
	r.Get("/getlistingbyid/{id}", s.Handler.HandleGetListingByID(storage))
	r.Post("/updatelistingbyid/{id}", s.Handler.HandleUpdateListingByID(storage))
	r.Delete("/deletelistingbyid/{id}", s.Handler.HandleDeleteListingByID(storage))

	r.Mount("/auth", s.AuthHandler.Routes())
	return r
}
