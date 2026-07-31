package team

import (
 "github.com/egor-shik/basketball-manager/internal/domain/coach"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
)

type Team struct {
 ID      int
 Name    string
 Coach   *coach.Coach
 Players []*player.Player
 Morale  TeamMorale
 Budget  Budget
}

func (t *Team) AddPlayer(p *player.Player) {
 t.Players = append(t.Players, p)
}

func (t *Team) RemovePlayer(playerID int) {
 for i, p := range t.Players {
  if p.Personal.ID == playerID {
   t.Players = append(t.Players[:i], t.Players[i+1:]...)
   break
  }
 }
}

func (t *Team) HasPlayer(p *player.Player) bool {
 for _, rosterPlayer := range t.Players {
  if rosterPlayer.Personal.ID == p.Personal.ID {
   return true
  }
 }
 return false
}

type Budget struct {
 Balance   int64
 SalaryCap int64
 Payroll   int64
}

func (b Budget) CurrentBalance() int64 {
 return b.Balance
}

func (b Budget) CanAfford(amount int64) bool {
 return b.Balance >= amount
}

func (b *Budget) AddPayroll(amount int64) {
 b.Payroll += amount
}

func (b *Budget) RemovePayroll(amount int64) {
 b.Payroll -= amount
 if b.Payroll < 0 {
  b.Payroll = 0
 }
}

func (b *Budget) AddExpense(amount int64) {
 b.Balance -= amount
}
