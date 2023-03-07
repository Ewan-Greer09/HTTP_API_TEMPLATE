package repository

import (
	"log"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

// creates a new record in the database
func (db *GormDatabase) CreateRecord(job *types.JobListing) {
	db.db.Create(&job)
}

// gets a record from the database
func (db *GormDatabase) GetRecord(jobID string) *types.JobListing {
	listing := &types.JobListing{}
	db.db.Where("id = ?", jobID).First(&listing)
	log.Println("Listing: ", listing)
	return listing
}

// TODO: add route to get all records
// gets all records from the database limit 100
func (db *GormDatabase) GetAllRecords() []*types.JobListing {
	listings := []*types.JobListing{}
	db.db.Limit(100).Find(&listings)
	return listings
}

// updates a record in the database
func (db *GormDatabase) UpdateRecord(job *types.JobListing) error {
	return nil
}

// deletes a record from the database
func (db *GormDatabase) DeleteRecord(jobID string) error {
	return nil
}
