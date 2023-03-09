package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/validationserver/validators"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

// HandleValidate is the handler for the /api/validate endpoint
// and passed data to the validator
func (h *Handler) HandleValidate(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleValidate called")
	jobListing := types.JobListing{}

	err := json.NewDecoder(r.Body).Decode(&jobListing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	violations := validators.ValidateJobBoardPostRequest(jobListing)
	log.Println(violations)
	if len(violations) > 0 {
		msg := fmt.Sprintf("validation failed: %v", violations)

		response := types.ApiResponse{
			Code: http.StatusBadRequest,
			Desc: msg,
			Data: violations,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}
