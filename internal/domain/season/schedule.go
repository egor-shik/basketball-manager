package season

import (
	"github.com/egor-shik/basketball-manager/internal/domain/match"
	"github.com/egor-shik/basketball-manager/internal/domain/team"
)

type Schedule struct {
	matches      []*match.Match
	currentIndex int
}

func NewSchedule(teams []*team.Team) *Schedule {
	var matches []*match.Match
	matchID := 1

	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			matches = append(matches, match.NewMatch(matchID, teams[i], teams[j]))
			matchID++
		}
	}

	return &Schedule{
		matches:      matches,
		currentIndex: 0,
	}
}

func (sch *Schedule) HasNext() bool {
	return sch.currentIndex < len(sch.matches)
}

func (sch *Schedule) Next() *match.Match {
	if !sch.HasNext() {
		return nil
	}
	m := sch.matches[sch.currentIndex]
	sch.currentIndex++
	return m
}
