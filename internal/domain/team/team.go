package team

import (
 "github.com/egor-shik/basketball-manager/internal/domain/coach"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
)

type Budget struct {
 Balance   int64 
 SalaryCap int64 
 Payroll   int64 
}

func (b *Budget) CanAfford(amount int64) bool {
 return b.Balance >= amount
}

func (b *Budget) AddExpense(amount int64) {
 b.Balance -= amount
}

func (b *Budget) ReceiveMoney(amount int64) {
 b.Balance += amount
}

func (b *Budget) CurrentBalance() int64 {
 return b.Balance
}

func (b *Budget) AddPayroll(amount int64) {
 b.Payroll += amount
}

func (b *Budget) RemovePayroll(amount int64) {
 b.Payroll -= amount
}

type Team struct {
 Name    string
 Players []*player.Player
 Coach   *coach.Coach
 Budget  Budget
 Morale  TeamMorale
}

func (t *Team) AddPlayer(p *player.Player) {
 t.Players = append(t.Players, p)
}

func (t *Team) RemovePlayer(p *player.Player) bool {
 for i, current := range t.Players {
  if current.Personal.ID == p.Personal.ID {
   t.Players = append(t.Players[:i], t.Players[i+1:]...)
   return true
  }
 }
 return false
}

func (t *Team) HasPlayer(p *player.Player) bool {
 for _, current := range t.Players {
  if current.Personal.ID == p.Personal.ID {
   return true
  }
 }
 return false
}

func (t *Team) PlayerCount() int {
 return len(t.Players)
}
