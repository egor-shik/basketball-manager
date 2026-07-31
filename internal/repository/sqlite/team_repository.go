package sqlite

import (
 "database/sql"

 "github.com/egor-shik/basketball-manager/internal/domain/team"
)

type TeamRepository struct {
 db *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
 return &TeamRepository{db: db}
}

func (r *TeamRepository) Save(t *team.Team) error {
 query := "INSERT INTO teams (name, balance, salary_cap, payroll) VALUES (?, ?, ?, ?)"
 _, err := r.db.Exec(query, t.Name, t.Budget.Balance, t.Budget.SalaryCap, t.Budget.Payroll)
 return err
}

func (r *TeamRepository) FindByID(id int) (*team.Team, error) {
 query := "SELECT name, balance, salary_cap, payroll FROM teams WHERE id = ?"
 row := r.db.QueryRow(query, id)

 var t team.Team
 err := row.Scan(&t.Name, &t.Budget.Balance, &t.Budget.SalaryCap, &t.Budget.Payroll)
 if err != nil {
  return nil, err
 }
 return &t, nil
}

func (r *TeamRepository) FindAll() ([]*team.Team, error) {
 query := "SELECT name, balance, salary_cap, payroll FROM teams"
 rows, err := r.db.Query(query)
 if err != nil {
  return nil, err
 }
 defer rows.Close()

 var teams []*team.Team
 for rows.Next() {
  var t team.Team
  if err := rows.Scan(&t.Name, &t.Budget.Balance, &t.Budget.SalaryCap, &t.Budget.Payroll); err != nil {
   return nil, err
  }
  teams = append(teams, &t)
 }
 return teams, nil
}

func (r *TeamRepository) Update(t *team.Team) error {
 query := "UPDATE teams SET balance=?, salary_cap=?, payroll=? WHERE name=?"
 _, err := r.db.Exec(query, t.Budget.Balance, t.Budget.SalaryCap, t.Budget.Payroll, t.Name)
 return err
}