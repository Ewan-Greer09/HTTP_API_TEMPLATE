package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/authserver/auth"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/authserver/redirect"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

func (h *Handler) HandleAuthenticate(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	//getlistingbyidRequestType := r.URL.Query().Get("getlistingbyid")


	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No Authorization Header", http.StatusUnauthorized)
		return
	}
	ok := auth.VerifyToken(authHeader)
	if !ok {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	authRequest := &types.Request{}
	err := json.NewDecoder(r.Body).Decode(authRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobListing := types.NewJobListingPerams(
		authRequest.JobListing.JobID,
		authRequest.JobListing.JobTitle,
		authRequest.JobListing.JobDescription,
		authRequest.JobListing.JobLocation,
		authRequest.JobListing.JobCompany,
		authRequest.JobListing.JobSalary,
	)

	err = redirect.RedirectRequestToAPI(jobListing, authRequest.RequestType, id)
	if err != nil {
		if err.Error() == "404 Not Found" {
			http.Error(w, "Invalid Request Type", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
