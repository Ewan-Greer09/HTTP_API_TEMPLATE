package repository

import (
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

// creates a new record in the database
func (db *SQLDatabase) CreateRecord(job *types.JobListing) error {
	return nil
}

// gets a record from the database
func (db *SQLDatabase) GetRecord(jobID string) (*types.JobListing, error) {
	return nil, nil
}

// updates a record in the database
func (db *SQLDatabase) UpdateRecord(job *types.JobListing) error {
	return nil
}

// deletes a record from the database
func (db *SQLDatabase) DeleteRecord(jobID string) error {
	return nil
}
