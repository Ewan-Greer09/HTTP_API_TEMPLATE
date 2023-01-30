package redirect

import (
	"context"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/carlmjohnson/requests"
)

// TODO: FIX THIS METHOD SO IT APPLIES TO ALL REQUEST TYPES
func RedirectRequestToAPI(request types.JobListing, requestType, id, getById string) error {
	ctx := context.Background()

	if getById != "" {
		err := requests.
			URL("http://localhost:8080/api/"+requestType+"/"+id).
			Method(http.MethodGet).
			Header("content-type", "application/json").
			CheckStatus(http.StatusOK).
			Fetch(ctx)

		if err != nil {
			return err
		}
	} else {
		err := requests.URL("http://localhost:8080/api/"+requestType).
			Method(http.MethodPost).
			Header("content-type", "application/json").
			CheckStatus(http.StatusOK).
			BodyJSON(&request).
			Fetch(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
