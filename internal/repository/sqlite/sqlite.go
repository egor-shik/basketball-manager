package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SQLite struct {
	db *sql.DB
}

// Opens a connection to the SQLite database using the provided DSN and verifies it with a ping
func New(dsn string) (*SQLite, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &SQLite{db: db}, nil
}

// DB returns the underlying *sql.DB instance for repository initialization
func (s *SQLite) DB() *sql.DB {
	return s.db
}

// Safely terminates the database connection pool
func (s *SQLite) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
