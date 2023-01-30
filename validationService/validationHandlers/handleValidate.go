package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/validationService/validators"
)

// HandleValidate is the handler for the /api/validate endpoint
// it parsed passed data to the validator
func HandleValidate(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleValidate called")
	jobListing := types.JobListing{}

	err := json.NewDecoder(r.Body).Decode(&jobListing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//pass jobListing to validator
	violations := validators.ValidateJobBoardCreateListing(jobListing)
	log.Println(violations)
	if len(violations) > 0 {
		err := fmt.Errorf("validation failed: %v", violations)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
