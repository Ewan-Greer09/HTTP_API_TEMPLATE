package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

func (h *Handler) HandleCreateListing(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newListing, err := h.CreateNewListing(r, storage)
		if err != nil {
			log.Println("ERROR: ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(newListing)
		if err != nil {
			log.Println("Error encoding response body")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
