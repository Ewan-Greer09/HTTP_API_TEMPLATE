package handlers

import (
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleUpdateListingByID(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := storage[id]; !ok {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		listing, err := h.UpdateJobListing(r, storage, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		storage[id] = listing

		h.logger.Info("Listing updated")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
