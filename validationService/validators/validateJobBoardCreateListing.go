package validators

import (
	"log"

	"gopkg.in/go-playground/validator.v9"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

type Violation struct {
	Field string
	Desc  string
}

func ValidateJobBoardCreateListing(jobListing types.JobListing) []Violation {
	var violations []Violation

	log.Println("ValidateJobBoardCreateListing called")

	validator := validator.New()

	// TODO: write an algorithm to validate the jobListing
	errs := validator.Var(jobListing.JobID, "required,numeric")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobID", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.JobTitle, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobTitle", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.JobDescription, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobDescription", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.JobLocation, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobLocation", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.JobSalary, "required,numeric")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobSalary", Desc: "missing field"})
	}

	errs = validator.Var(jobListing.JobCompany, "required")
	if errs != nil {
		violations = append(violations, Violation{Field: "JobCompany", Desc: "missing field"})
	}

	return violations
}
