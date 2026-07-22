package sqlite

import (
    "database/sql"
     "github.com/egor-shik/basketball-manager/internal/domain/news"
)

type NewsRepository struct {
    db *sql.DB
}

func NewNewsRepository(db *sql.DB) *NewsRepository {
    return &NewsRepository{db: db}
}

func (r *NewsRepository) Save(n *news.News) error {
    return nil
}

func (r *NewsRepository) FindAll() ([]news.News, error) {
    return nil, nil
}