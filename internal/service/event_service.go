package service

import (
	"time"

	"github.com/egor-shik/basketball-manager/internal/domain/event"
	"github.com/egor-shik/basketball-manager/internal/domain/match"
)

// EventService handles the creation and initialization of domain events
type EventService struct {
	currentID int
}

// Instantiates a new EventService with an initial ID counter
func NewEventService() *EventService {
	return &EventService{currentID: 1}
}

// nextID safely increments and returns the next unique event identifier
func (s *EventService) nextID() int {
	id := s.currentID
	s.currentID++
	return id
}

// Constructs an event record for a completed match
func (s *EventService) CreateMatchFinishedEvent(seasonNum int, m *match.Match, res *match.MatchResult) event.Event {
	return event.Event{
		ID:           s.nextID(),
		Type:         event.MatchFinished,
		MatchID:      m.ID,
		SeasonNumber: seasonNum,
		Timestamp:    time.Now(),
		TeamName:     m.HomeTeam.Name,
		OpponentName: m.AwayTeam.Name,
		HomeScore:    res.HomeScore,
		AwayScore:    res.AwayScore,
	}
}

// Constructs an event record for a player acquisition
func (s *EventService) CreatePlayerSignedEvent(teamName string, playerName string) event.Event {
	return event.Event{
		ID:         s.nextID(),
		Type:       event.PlayerSigned,
		Timestamp:  time.Now(),
		TeamName:   teamName,
		PlayerName: playerName,
	}
}

// Constructs an event record for a player release
func (s *EventService) CreatePlayerReleasedEvent(teamName string, playerName string) event.Event {
	return event.Event{
		ID:         s.nextID(),
		Type:       event.PlayerReleased,
		Timestamp:  time.Now(),
		TeamName:   teamName,
		PlayerName: playerName,
	}
}

// Constructs an event record for a completed trade between two teams
func (s *EventService) CreateTradeEvent(teamAName, teamBName, playerAName string) event.Event {
	return event.Event{
		ID:           s.nextID(),
		Type:         event.TradeCompleted,
		Timestamp:    time.Now(),
		TeamName:     teamAName,
		OpponentName: teamBName,
		PlayerName:   playerAName,
	}
}

// Constructs an event record for the conclusion of a season
func (s *EventService) CreateSeasonFinishedEvent(seasonNum int, championName string) event.Event {
	return event.Event{
		ID:           s.nextID(),
		Type:         event.SeasonFinished,
		SeasonNumber: seasonNum,
		Timestamp:    time.Now(),
		TeamName:     championName,
	}
}
