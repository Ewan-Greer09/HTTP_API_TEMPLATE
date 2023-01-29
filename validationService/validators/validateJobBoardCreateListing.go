package validators

import (
	"gopkg.in/go-playground/validator.v9"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

type Violation struct {
	Field string
	Desc  string
}

func ValidateJobBoardCreateListing(jobListing types.JobListing) []Violation {
	var violations []Violation

	validator := validator.New()

	err := validator.Struct(jobListing)
	if err != nil {
		violations = append(violations, Violation{Field: "Title", Desc: "Title is required"})
		return violations
	} else {
		return nil
	}
}
