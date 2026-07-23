package sqlite

import (
"database/sql"
"os"

_ "modernc.org/sqlite"
)

type SQLite struct {
db *sql.DB
}

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

func (s *SQLite) DB() *sql.DB {
return s.db
}

func (s *SQLite) Close() error {
if s.db != nil {
return s.db.Close()
}
return nil
}

func (s *SQLite) InitSchema(filepath ...string) error {
if len(filepath) > 0 && filepath[0] != "" {
content, err := os.ReadFile(filepath[0])
if err == nil {
_, err = s.db.Exec(string(content))
return err
}
}

query := `
CREATE TABLE IF NOT EXISTS players (
id TEXT PRIMARY KEY,
name TEXT NOT NULL,
position TEXT NOT NULL,
skill INTEGER NOT NULL,
team_id TEXT
);

CREATE TABLE IF NOT EXISTS teams (
id TEXT PRIMARY KEY,
name TEXT NOT NULL,
budget REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS events (
id TEXT PRIMARY KEY,
type TEXT NOT NULL,
description TEXT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS news (
id TEXT PRIMARY KEY,
title TEXT NOT NULL,
content TEXT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

_, err := s.db.Exec(query)
return err
}
