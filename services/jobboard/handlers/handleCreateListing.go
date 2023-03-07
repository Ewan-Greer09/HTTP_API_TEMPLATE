package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
)

func (h *Handler) HandleCreateListing(db *repository.GormDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newListing, err := h.CreateNewListing(r, db)
		if err != nil {
			h.logger.Error("Error creating new listing")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(newListing)
		if err != nil {
			h.logger.Error("Error encoding new listing")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
