package handlers

import (
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleDeleteListingByID(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, ok := storage[id]; !ok {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		delete(storage, id)
		h.logger.Info("Listing deleted")

		w.WriteHeader(http.StatusOK)
	}
}
