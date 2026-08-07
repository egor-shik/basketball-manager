package news

import "time"

type Article struct {
	ID             int
	Title          string
	Body           string
	CreatedAt      time.Time
	RelatedEventID int
}
