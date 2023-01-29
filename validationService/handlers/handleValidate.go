package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/validationService/validators"
)

// HandleValidate is the handler for the /api/validate endpoint
// and passed data to the validator
func HandleValidate(w http.ResponseWriter, r *http.Request) {
	jobListing := types.JobListing{}

	err := json.NewDecoder(r.Body).Decode(&jobListing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//pass jobListing to validator
	violations := validators.ValidateJobBoardCreateListing(jobListing)
	if len(violations) > 0 {
		http.Error(w, violations[0].Desc, http.StatusBadRequest)
		return
	}
}
