package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/client"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

type Handler struct {
	HandlerInterface HandlerInterface
}

type HandlerInterface interface {
	HandleValidateRequest(listing *types.JobListing) error
}

func NewHandler() *Handler {
	return &Handler{}
}

// TODO: rework so that propper error message is returned to the user
// ! currently returns an error from decdoing but want to return an error from the validator
// CreateNewListing creates a new listing from the request body and adds it to the storage
func (h *Handler) CreateNewListing(r *http.Request, storage map[string]types.JobListing) (*types.JobListing, error) {
	newListing := types.NewJobListing()
	err := json.NewDecoder(r.Body).Decode(&newListing)
	if err != nil {
		log.Println("Error decoding request body")
		return nil, err
	}

	uuid := uuid.New()
	newListing.ID = uuid.String()

	err = handleValidateRequest(&newListing)
	if err != nil {
		log.Println("Error validating request body: " + err.Error())
		errstr := "Error validating request body: Code: " + strconv.FormatInt(400, 10)
		return nil, errors.New(errstr)
	}

	log.Println("Created new listing: \n", spew.Sdump(newListing))
	storage[newListing.ID] = newListing

	return &newListing, nil
}

func handleValidateRequest(listing *types.JobListing) error {
	c := client.NewClient()
	ok, err := c.SendValidateRequest(listing)
	if !ok {
		return err
	}

	return nil
}
