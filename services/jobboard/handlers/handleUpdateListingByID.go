package handlers

import (
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
)

func (h *Handler) HandleUpdateListingByID(db *repository.SQLDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if listing := db.GetRecord(id); listing == nil {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}

		listing, err := h.UpdateJobListing(r, db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db.CreateRecord(&listing)

		log.Println("Updated listing: \n", spew.Sdump(listing))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
