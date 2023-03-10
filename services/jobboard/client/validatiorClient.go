package client

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	types "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/jobboard"
	validation "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/validationSever"
	"github.com/carlmjohnson/requests"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

// SendValidateRequest sends a request to the validation service to validate the listing
func (c *Client) SendValidateRequest(listing *types.JobListing) (*validation.ApiResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl := NewClientWithRetries(3, 10*time.Second)
	response := &validation.ApiResponse{}

	err := requests.URL("http://localhost:3000/api/validate").
		Client(cl).
		Method(http.MethodPost).
		Header("content-type", "application/json").
		CheckStatus(http.StatusOK).
		BodyJSON(&listing).
		ToJSON(response).
		Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// NewClientWithRetries returns a new http client with retry logic
func NewClientWithRetries(count int, maxDuration time.Duration) *http.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = count
	client.RetryWaitMax = maxDuration
	client.RetryWaitMin = 1 * time.Second

	return client.StandardClient()
}
