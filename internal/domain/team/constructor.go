package team

import (
	"github.com/egor-shik/basketball-manager/internal/domain/coach"
	"github.com/egor-shik/basketball-manager/internal/domain/player"
)

// It enforces slice allocation to prevent nil panics during roster operations
// and initializes default morale and chemistry values
func NewTeam(id int, name string, c *coach.Coach, players []*player.Player) *Team {
	if players == nil {
		players = make([]*player.Player, 0)
	}
	return &Team{
		ID:      id,
		Name:    name,
		Coach:   c,
		Players: players,
		Morale: TeamMorale{
			TeamMorale: 100.0,
			Chemistry:  1.0,
		},
	}
}
