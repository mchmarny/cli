package data

import (
	"database/sql"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const (
	dataFile string = "data.db"
)

type Service struct {
	db *sql.DB
	mu sync.Mutex
}

// New initializes the database for a given name.
func New(name string) (*Service, error) {
	dataDir, created, err := getOrCreateHomeDir(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get or create home dir")
	}

	s := &Service{}

	dataPath := filepath.Join(dataDir, dataFile)
	s.db, err = sql.Open("sqlite3", dataPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database: %s", dataPath)
	}

	if created {
		if _, err := s.db.Exec(ddl); err != nil {
			return nil, errors.Wrapf(err, "failed to create database schema in: %s", dataPath)
		}
	}

	return s, nil
}

// Close closes the database if one of previously created.
func (s *Service) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
