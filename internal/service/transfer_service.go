package service

import (
 "github.com/egor-shik/basketball-manager/internal/domain/player"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
)

type TransferService struct{}

func NewTransferService() *TransferService {
 return &TransferService{}
}

func (s *TransferService) SignPlayer(t *team.Team, p *player.Player) error {
 if t.HasPlayer(p) {
  return ErrPlayerAlreadySigned
 }
 if !t.Budget.CanAfford(p.Contract.Salary) {
  return ErrNotEnoughMoney
 }
 if t.Budget.Payroll+p.Contract.Salary > t.Budget.SalaryCap {
  return ErrSalaryCapExceeded
 }

 t.AddPlayer(p)
 t.Budget.AddPayroll(p.Contract.Salary)
 t.Budget.AddExpense(p.Contract.Salary)

 return nil
}

func (s *TransferService) ReleasePlayer(t *team.Team, p *player.Player) error {
 if !t.HasPlayer(p) {
  return ErrPlayerNotFound
 }

 t.RemovePlayer(p.Personal.ID)
 t.Budget.RemovePayroll(p.Contract.Salary)

 return nil
}

func (s *TransferService) TradePlayers(teamA, teamB *team.Team, playerA, playerB *player.Player) error {
 if !teamA.HasPlayer(playerA) || !teamB.HasPlayer(playerB) {
  return ErrPlayerNotFound
 }

 projectedPayrollA := teamA.Budget.Payroll - playerA.Contract.Salary + playerB.Contract.Salary
 if projectedPayrollA > teamA.Budget.SalaryCap {
  return ErrSalaryCapExceeded
 }

 projectedPayrollB := teamB.Budget.Payroll - playerB.Contract.Salary + playerA.Contract.Salary
 if projectedPayrollB > teamB.Budget.SalaryCap {
  return ErrSalaryCapExceeded
 }

 teamA.Budget.RemovePayroll(playerA.Contract.Salary)
 teamA.Budget.AddPayroll(playerB.Contract.Salary)

 teamB.Budget.RemovePayroll(playerB.Contract.Salary)
 teamB.Budget.AddPayroll(playerA.Contract.Salary)

 teamA.RemovePlayer(playerA.Personal.ID)
 teamA.AddPlayer(playerB)

 teamB.RemovePlayer(playerB.Personal.ID)
 teamB.AddPlayer(playerA)

 return nil
}
