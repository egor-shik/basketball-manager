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
    return nil
}

func (r *TeamRepository) FindByID(id int) (*team.Team, error) {
    return nil, nil
}

func (r *TeamRepository) FindAll() ([]*team.Team, error) {
    return nil, nil
}