package client

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/carlmjohnson/requests"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

// SendValidateRequest sends a request to the validation service to validate the listing
func (c *Client) SendValidateRequest(listing *types.JobListing) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := NewClientWithRetries(3, 10*time.Second)

	err := requests.URL("http://localhost:3000/api/validate").
		Client(client).
		Method(http.MethodPost).
		Header("content-type", "application/json").
		CheckStatus(http.StatusOK).
		BodyJSON(&listing).
		Fetch(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

// NewClientWithRetries returns a new http client with retry logic
func NewClientWithRetries(count int, maxDuration time.Duration) *http.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = count
	client.RetryWaitMax = maxDuration
	client.RetryWaitMin = 1 * time.Second

	return client.StandardClient()
}
