package sqlite

import (
	"database/sql"

	"github.com/egor-shik/basketball-manager/internal/domain/player"
)

type PlayerRepository struct {
	db *sql.DB
}

// Instantiates a new SQLite-backed PlayerRepository
func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Inserts a player entity into the database and assigns the generated ID
func (r *PlayerRepository) Save(p *player.Player) error {
	query := "INSERT INTO players (first_name, last_name, age, position, salary, market_value) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, p.Personal.FirstName, p.Personal.LastName, p.Personal.Age, string(p.Personal.Pos), p.Contract.Salary, p.Contract.MarketValue)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		p.Personal.ID = int(id)
	}
	return nil
}

// Queries and returns a player by their unique identifier
func (r *PlayerRepository) FindByID(id int) (*player.Player, error) {
	query := "SELECT id, first_name, last_name, age, position, salary, market_value FROM players WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var p player.Player
	var pos string
	err := row.Scan(&p.Personal.ID, &p.Personal.FirstName, &p.Personal.LastName, &p.Personal.Age, &pos, &p.Contract.Salary, &p.Contract.MarketValue)
	if err != nil {
		return nil, err
	}
	p.Personal.Pos = player.Position(pos)
	return &p, nil
}

// Fetches all player records from the database
func (r *PlayerRepository) FindAll() ([]*player.Player, error) {
	query := "SELECT id, first_name, last_name, age, position, salary, market_value FROM players"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []*player.Player
	for rows.Next() {
		var p player.Player
		var pos string
		if err := rows.Scan(&p.Personal.ID, &p.Personal.FirstName, &p.Personal.LastName, &p.Personal.Age, &pos, &p.Contract.Salary, &p.Contract.MarketValue); err != nil {
			return nil, err
		}
		p.Personal.Pos = player.Position(pos)
		players = append(players, &p)
	}
	return players, nil
}

// Modifies existing player records in the database
func (r *PlayerRepository) Update(p *player.Player) error {
	query := "UPDATE players SET first_name=?, last_name=?, age=?, position=?, salary=?, market_value=? WHERE id=?"
	_, err := r.db.Exec(query, p.Personal.FirstName, p.Personal.LastName, p.Personal.Age, string(p.Personal.Pos), p.Contract.Salary, p.Contract.MarketValue, p.Personal.ID)
	return err
}
