package handlers

import (
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleDeleteListingByID(db *repository.SQLDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if _, err := db.GetRecord(id); err != nil {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		db.DeleteRecord(id)

		w.WriteHeader(http.StatusOK)
	}
}
