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

func (r *NewsRepository) Save(n *news.Article) error {
 query := "INSERT INTO news (title, body, created_at, related_event_id) VALUES (?, ?, ?, ?)"
 res, err := r.db.Exec(query, n.Title, n.Body, n.CreatedAt, n.RelatedEventID)
 if err != nil {
  return err
 }
 id, err := res.LastInsertId()
 if err == nil {
  n.ID = int(id)
 }
 return nil
}

func (r *NewsRepository) FindAll() ([]news.Article, error) {
 query := "SELECT id, title, body, created_at, related_event_id FROM news"
 rows, err := r.db.Query(query)
 if err != nil {
  return nil, err
 }
 defer rows.Close()

 var articles []news.Article
 for rows.Next() {
  var n news.Article
  if err := rows.Scan(&n.ID, &n.Title, &n.Body, &n.CreatedAt, &n.RelatedEventID); err != nil {
   return nil, err
  }
  articles = append(articles, n)
 }
 return articles, nil
}