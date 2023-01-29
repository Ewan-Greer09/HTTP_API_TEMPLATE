package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/jobBoardService/handlers"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/storage"
)

func StartServer() {
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Starting HTTP API Template Service...")
	log.Println("Populating storage...")
	storage := storage.PopulateStorage()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Get("/ping", handlers.HandlePing)
		r.Post("/createlisting", handlers.HandleCreateListing(storage))
		r.Get("/getlistingbyid/{id}", handlers.HandleGetListingByID(storage))
		r.Delete("/deletelistingbyid/{id}", handlers.HandleDeleteListingByID(storage))
		r.Put("/updatelistingbyid/{id}", handlers.HandleUpdateListingByID(storage))
	})

	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Listening and serving on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
