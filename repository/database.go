package repository

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/logger"
	"github.com/Ewan-Greer09/HTTP_API_TEMPLATE/types"
)

var (
	ErrDuplicate    = errors.New("duplicate entry")
	ErrNotExist     = errors.New("not found")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// This is a database repository
type SQLDatabase struct {
	db *sql.DB
}

// Allows for crud operations on the db
type DatabaseRepository interface {
	CreateRecord(*types.JobListing) error
	GetRecord(string) (*types.JobListing, error)
	UpdateRecord(*types.JobListing) error
	DeleteRecord(string) error
}

func NewDatabase(logger *logger.Logger) (*SQLDatabase, error) {
	os.Remove("./database.db")

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		logger.Errorf("Failed to open database: %v", err)
		return nil, err
	}
	defer db.Close()

	sqlStmt := `create table joblisting (id text not null primary key, position text, description text, location text, pay real, company text, salaried bool, remote bool, datafields text);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		logger.Errorf("Failed to create table: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &SQLDatabase{
		db: db,
	}, nil
}
