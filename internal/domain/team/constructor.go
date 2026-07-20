package team

import (
 "github.com/egor-shik/basketball-manager/internal/domain/coach"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
)

func NewTeam(name string, c *coach.Coach, players []*player.Player) *Team {
 return &Team{
  Name:    name,
  Players: players,
  Coach:   c,
  Budget:  Budget{},
  Morale:  TeamMorale{},
 }
}