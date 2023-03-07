package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleGetListingByID(db *repository.SQLDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// *lookups to be done in seperate function
		listing, err := db.GetRecord(id)
		if err != nil {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		parsedListing, err := json.Marshal(listing)
		if err != nil {
			log.Println("Error parsing listing")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//write a response as a json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(parsedListing)
	}
}
