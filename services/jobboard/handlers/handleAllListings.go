package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"
)

func (h *Handler) HandleAllListings(db *repository.GormDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		listings := db.GetAllRecords()

		response, err := json.Marshal(listings)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
