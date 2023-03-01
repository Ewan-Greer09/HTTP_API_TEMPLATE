package handlers

import (
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/client"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

type Handler struct {
	HandlerInterface HandlerInterface
}

type HandlerInterface interface {
	HandleValidateRequest(listing *types.JobListing) error
}

func NewHandler() *Handler {
	return &Handler{
		HandlerInterface: &Handler{},
	}
}

func handleValidateRequest(listing *types.JobListing) error {
	c := client.NewClient()
	ok, err := c.SendValidateRequest(listing)
	if !ok {
		return err
	}

	return nil
}
