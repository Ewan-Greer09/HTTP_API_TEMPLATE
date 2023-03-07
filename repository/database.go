package repository

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/repository/models"
)

var (
	ErrDuplicate    = errors.New("duplicate entry")
	ErrNotExist     = errors.New("not found")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// This is a database repository
type SQLDatabase struct {
	db *gorm.DB
}

func NewDatabase(logger *logger.Logger) (*SQLDatabase, error) {
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

	return &SQLDatabase{db: db}, nil
}
