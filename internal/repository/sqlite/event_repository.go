package sqlite

import (
    "database/sql"
    "github.com/egor-shik/basketball-manager/internal/domain/event"
)

type EventRepository struct {
    db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
    return &EventRepository{db: db}
}

func (r *EventRepository) Save(e *event.Event) error {
    return nil
}

func (r *EventRepository) FindAll() ([]event.Event, error) {
    return nil, nil
}