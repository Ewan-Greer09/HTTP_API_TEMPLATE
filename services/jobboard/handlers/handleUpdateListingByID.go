package handlers

import (
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/davecgh/go-spew/spew"

	"github.com/go-chi/chi"
)

func (h *Handler) UpdateJobListing(db *repository.GormDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		listing := db.GetRecord(id)
		if listing == nil {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		db.UpdateRecord(listing)

		log.Println("Updated listing: \n", spew.Sdump(listing))

		w.WriteHeader(http.StatusOK)
	}
}
