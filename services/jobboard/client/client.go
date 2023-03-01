package client

import (
	"context"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/carlmjohnson/requests"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendValidateRequest(listing *types.JobListing) (bool, error) {
	ctx := context.Background()

	err := requests.URL("http://localhost:3000/api/validate").
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
