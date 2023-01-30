package server

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/validationService/validationHandlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartValidationServer() {
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Starting HTTP API Template Service...")
	log.Println("Populating storage...")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/api", func(r chi.Router) {
		r.Post("/validate", handlers.HandleValidate)
	})

	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "Listening and serving on port :8080")
	log.Fatal(http.ListenAndServe(":8085", router))
}
