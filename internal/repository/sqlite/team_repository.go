package sqlite

import (
	"database/sql"

	"github.com/egor-shik/basketball-manager/internal/domain/team"
)

// Implements repository.TeamRepository for SQLite storage
type TeamRepository struct {
	db *sql.DB
}

// Instantiates a new SQLite-backed TeamRepository
func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// Inserts a team record into the database.
func (r *TeamRepository) Save(t *team.Team) error {
	query := "INSERT INTO teams (name, balance, salary_cap, payroll) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, t.Name, t.Budget.Balance, t.Budget.SalaryCap, t.Budget.Payroll)
	return err
}

// Queries and returns a team by its unique identifier
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

// Retrieves all team records from the database
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

// WARNING: Currently updates by team name. Consider switching to primary key (ID) if duplicate names are possible
func (r *TeamRepository) Update(t *team.Team) error {
	query := "UPDATE teams SET balance=?, salary_cap=?, payroll=? WHERE name=?"
	_, err := r.db.Exec(query, t.Budget.Balance, t.Budget.SalaryCap, t.Budget.Payroll, t.Name)
	return err
}
