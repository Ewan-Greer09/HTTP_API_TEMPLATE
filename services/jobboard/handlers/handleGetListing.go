package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleGetListingByID(db *repository.GormDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		listing := db.GetRecord(id)
		if listing == nil {
			http.Error(w, "Listing does not exist", http.StatusNotFound)
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
