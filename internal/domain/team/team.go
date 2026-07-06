package team

import (
	"github.com/egor-shik/basketball-manager/internal/domain/coach"
	"github.com/egor-shik/basketball-manager/internal/domain/player"
)

type Team struct {
	Name    string
	Players []*player.Player
	Coach   *coach.Coach
	Budget  Budget
	Morale  TeamMorale
}

type Budget struct {
	Balance   int
	SalaryCap int
	Payroll   int
}
