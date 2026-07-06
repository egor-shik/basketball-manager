package match

import (
	"github.com/egor-shik/basketball-manager/internal/domain/team"
)

type MatchStatus int

const (
	StatusNotStarted MatchStatus = iota
	StatusFinished
)

type Match struct {
	ID       int
	HomeTeam *team.Team
	AwayTeam *team.Team
	Result   *MatchResult
	Status   MatchStatus
}

func NewMatch(id int, home, away *team.Team) *Match {
	return &Match{
		ID:       id,
		HomeTeam: home,
		AwayTeam: away,
		Status:   StatusNotStarted,
	}
}
