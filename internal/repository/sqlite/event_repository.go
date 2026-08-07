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
	query := "INSERT INTO events (type, timestamp) VALUES (?, ?)"
	res, err := r.db.Exec(query, string(e.Type), e.Timestamp)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		e.ID = int(id)
	}
	return nil
}

func (r *EventRepository) FindAll() ([]event.Event, error) {
	query := "SELECT id, type, timestamp FROM events"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []event.Event
	for rows.Next() {
		var e event.Event
		var eventType string
		if err := rows.Scan(&e.ID, &eventType, &e.Timestamp); err != nil {
			return nil, err
		}
		e.Type = event.Type(eventType)
		events = append(events, e)
	}
	return events, nil
}
