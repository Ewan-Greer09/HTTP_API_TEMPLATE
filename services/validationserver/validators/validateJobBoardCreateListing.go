package validators

import (
	"log"

	"gopkg.in/go-playground/validator.v9"

	jobboard "github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types/jobboard"
)

type Violation struct {
	Field string
	Desc  string
}

func ValidateJobBoardPostRequest(jobListing jobboard.JobListing) []Violation {
	var violations []Violation

	log.Println("ValidateJobBoardPostRequest called")

	validator := validator.New()

	// TODO: write an algorithm to validate the jobListing
	// TODO: update function to reflect new types.JobListing struct
	errs := validator.Var(jobListing.Position, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobTitle", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.Description, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobDescription", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.Location, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobLocation", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.Pay, "required,numeric")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobSalary", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.Company, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobCompany", Desc: "missing field"})
	}

	return violations
}
