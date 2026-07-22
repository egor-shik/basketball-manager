package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
    db *sql.DB
}

func New(path string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &SQLite{db; db}, nil
}