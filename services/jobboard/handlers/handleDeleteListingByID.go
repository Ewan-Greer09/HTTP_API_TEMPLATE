package handlers

import (
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleDeleteListingByID(db *repository.GormDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		listing := db.GetRecord(id)
		if listing == nil {
			http.Error(w, "Listing does not exist", http.StatusNotFound)
			return
		}
    
		db.DeleteRecord(id)

		w.WriteHeader(http.StatusOK)
	}
}
