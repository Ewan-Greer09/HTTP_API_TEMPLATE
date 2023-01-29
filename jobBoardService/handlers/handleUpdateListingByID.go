package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/go-chi/chi"
)

func HandleUpdateListingByID(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := storage[id]; !ok {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		newListing := types.NewJobListing()
		err := json.NewDecoder(r.Body).Decode(&newListing)
		if err != nil {
			log.Println("Error decoding request body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		storage[id] = newListing
		w.WriteHeader(http.StatusOK)
	}
}
