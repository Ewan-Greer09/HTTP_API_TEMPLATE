package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

func (h *Handler) HandleCreateListing(storage map[string]types.JobListing) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newListing := types.NewJobListing()
		err := json.NewDecoder(r.Body).Decode(&newListing)
		if err != nil {
			log.Println("Error decoding request body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate a new UUID for the new listing
		uuid := uuid.New()
		newListing.ID = uuid.String()

		//err = h.HandleValidateRequest(&newListing)
		//if err != nil {
		//	log.Println("Error validating request body")
		//	errstr := "Error validating request body: Code: " + strconv.FormatInt(400, 10)
		//	w.Write([]byte(errstr))
		//	return
		//}

		storage[newListing.ID] = newListing
		log.Println("Created new listing: \n", spew.Sdump(newListing))

		w.WriteHeader(http.StatusCreated)
	}
}
