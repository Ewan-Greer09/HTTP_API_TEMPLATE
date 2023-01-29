package storage

import (
	"log"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

// *TEMP storage for job listings
var storage = make(map[string]types.JobListing)

type Storage struct {
	storage map[string]types.JobListing
}

func PopulateStorage() map[string]types.JobListing {
	storage = map[string]types.JobListing{
		"123": {
			JobID:          "123",
			JobTitle:       "Software Engineer",
			JobDescription: "Write code",
			JobLocation:    "London",
			JobSalary:      100000,
			JobCompany:     "Google",
		},
		"456": {
			JobID:          "456",
			JobTitle:       "Software Engineer",
			JobDescription: "Write code",
			JobLocation:    "London",
			JobSalary:      100000,
			JobCompany:     "Google",
		},
		"789": {
			JobID:          "789",
			JobTitle:       "Software Engineer",
			JobDescription: "Write code",
			JobLocation:    "London",
			JobSalary:      100000,
			JobCompany:     "Google",
		},
	}
	log.Println("Populated storage")
	return storage
}
