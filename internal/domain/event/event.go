package event

import (
	"time"
)

type Event struct {
	ID           int
	Type         Type
	MatchID      int
	SeasonNumber int
	Timestamp    time.Time

	PlayerID     int
	PlayerName   string
	TeamName     string
	OpponentName string
	HomeScore    int
	AwayScore    int
}
