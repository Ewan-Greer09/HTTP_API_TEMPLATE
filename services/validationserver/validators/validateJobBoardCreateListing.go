package validators

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	jobboard "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/jobboard"
)

type Violation struct {
	Desc string
}

// Validates struct contains all required fields using ozzo-validation
func ValidateJobBoardPostRequest(jobListing jobboard.JobListing) []Violation {
	var violations []Violation

	log.Println("ValidateJobBoardPostRequest called")

	err := validation.ValidateStruct(&jobListing,
		validation.Field(&jobListing.ID, validation.Required, is.UUID),
		validation.Field(&jobListing.Position, validation.Required, is.Alphanumeric),
		validation.Field(&jobListing.Description, validation.Required),
		validation.Field(&jobListing.Location, validation.Required),
		validation.Field(&jobListing.Pay, validation.Required, is.Float),
		validation.Field(&jobListing.Company, validation.Required),
		validation.Field(&jobListing.Salaried, validation.Required),
		validation.Field(&jobListing.Remote, validation.Required),
	)

	if err != nil {
		for _, err := range err.(validation.Errors) {
			violations = append(violations, Violation{
				Desc: err.Error(),
			})
		}
	}

	return violations
}
