package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/client"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/config"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

type Handler struct {
	HandlerInterface HandlerInterface
	cfg              config.JobBoardConfig
	logger           *logger.Logger
}

type HandlerInterface interface {
	HandleValidateRequest(listing *types.JobListing) error
}

func NewHandler(cfg config.JobBoardConfig, logger *logger.Logger) *Handler {
	return &Handler{
		cfg:    cfg,
		logger: logger,
	}
}

// CreateNewListing creates a new listing from the request body and adds it to the storage
func (h *Handler) CreateNewListing(r *http.Request, db *repository.GormDatabase) (*types.JobListing, error) {
	newListing := types.NewJobListing()
	err := json.NewDecoder(r.Body).Decode(&newListing)
	if err != nil {
		h.logger.Error("Error decoding request body")
		return nil, err
	}

	uuid := uuid.New()
	newListing.ID = uuid.String()

	err = handleValidateRequest(&newListing)
	if err != nil {
		h.logger.Errorf("error validating request body: %s", err.Error())
		return nil, errors.New(fmt.Sprintf("error validating request body: %s", err.Error()))
	}

	log.Println("Created new listing: \n", spew.Sdump(newListing))
	db.CreateRecord(&newListing)

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
