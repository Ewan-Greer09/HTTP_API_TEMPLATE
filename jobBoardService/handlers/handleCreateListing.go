package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

func (h *Handler) HandleCreateListing(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newListing := types.NewJobListing()
		err := json.NewDecoder(r.Body).Decode(&newListing)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//TODO: db call
		storage[newListing.ID] = newListing

		w.WriteHeader(http.StatusOK)
	}
}
