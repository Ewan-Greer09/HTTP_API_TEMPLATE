package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/go-chi/chi"
)

func HandleGetListingByID(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// *lookups to be done in seperate function
		listing, ok := storage[id]
		if !ok {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		parsedListing, err := json.Marshal(listing)
		if err != nil {
			log.Println("Error parsing listing")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(parsedListing)
	}
}
