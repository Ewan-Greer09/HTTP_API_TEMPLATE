package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository/models"
)

type GormDatabase struct {
	db *gorm.DB
}

func NewDatabase(logger *logger.Logger) (*GormDatabase, error) {
	// using default config for now
	db, err := gorm.Open(sqlite.Open("job_listings.db"), &gorm.Config{})
	if err != nil {
		logger.Panic(err)
	}

	db.Model(&models.JobListing{})

	// Migrate the schema
	err = db.AutoMigrate(&models.JobListing{})
	if err != nil {
		logger.Panic(err)
	}

	return &GormDatabase{db: db}, nil
}
