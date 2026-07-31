package sqlite

import (
	"database/sql"
	"fmt"
	"os"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) RunFromFile(filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %w", filepath, err)
	}

	if _, err := m.db.Exec(string(content)); err != nil {
		return fmt.Errorf("failed to execute migration %s: %w", filepath, err)
	}
	return nil
}
