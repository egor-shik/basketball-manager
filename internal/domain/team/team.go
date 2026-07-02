package team

import (
    "basketball-sim/internal/domain/player"
    "basketball-sim/internal/domain/coach"
)

type Team struct {
    Name    string
    Players []*player.Player
    Coach   *coach.Coach
    Budget  Budget
    Morale  TeamMorale
}

type Budget struct {
Balance int 
SalaryCap int
Payroll int
}