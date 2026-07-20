package service

import (
 "time"

 "github.com/egor-shik/basketball-manager/internal/domain/event"
 "github.com/egor-shik/basketball-manager/internal/domain/match"
)

type EventService struct {
 currentID int
}

func NewEventService() *EventService {
 return &EventService{currentID: 1}
}

func (s *EventService) nextID() int {
 id := s.currentID
 s.currentID++
 return id
}

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

func (s *EventService) CreatePlayerSignedEvent(teamName string, playerName string) event.Event {
	return event.Event{
	 ID:         s.nextID(),
	 Type:       event.PlayerSigned,
	 Timestamp:  time.Now(),
	 TeamName:   teamName,
	 PlayerName: playerName,
	}
   }
   
   func (s *EventService) CreatePlayerReleasedEvent(teamName string, playerName string) event.Event {
	return event.Event{
	 ID:         s.nextID(),
	 Type:       event.PlayerReleased,
	 Timestamp:  time.Now(),
	 TeamName:   teamName,
	 PlayerName: playerName,
	}
   }
   
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
   
   func (s *EventService) CreateSeasonFinishedEvent(seasonNum int, championName string) event.Event {
	return event.Event{
	 ID:           s.nextID(),
	 Type:         event.SeasonFinished,
	 SeasonNumber: seasonNum,
	 Timestamp:    time.Now(),
	 TeamName:     championName,
	}
   }
   