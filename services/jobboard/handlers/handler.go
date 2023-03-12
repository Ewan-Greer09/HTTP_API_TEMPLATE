package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/client"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/services/jobboard/config"
	types "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/jobboard"
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

	log.Println("Created new listing: \n", spew.Sdump(newListing))
	db.CreateRecord(&newListing)

	return &newListing, nil
}

func handleValidateRequest(listing *types.JobListing) (string, error) {
	c := client.NewClient()
	resp, err := c.SendValidateRequest(listing)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
