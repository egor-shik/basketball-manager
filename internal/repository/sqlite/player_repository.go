package sqlite

import (
    "database/sql"
    "github.com/egor-shik/basketball-manager/internal/domain/player"
)

type PlayerRepository struct {
    db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
    return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Save(p *player.Player) error {
    return nil
}

func (r *PlayerRepository) FindByID(id int) (*player.Player, error) {
    return nil, nil
}

func (r *PlayerRepository) FindAll() ([]*player.Player, error) {
    return nil, nil
}

func (r *PlayerRepository) Update(p *player.Player) error {
    return nil
}