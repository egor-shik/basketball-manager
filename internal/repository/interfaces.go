package repository

import (
 "github.com/egor-shik/basketball-manager/internal/domain/event"
 "github.com/egor-shik/basketball-manager/internal/domain/news"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
)

type PlayerRepository interface {
 Save(p *player.Player) error
 FindByID(id int) (*player.Player, error)
 FindAll() ([]*player.Player, error)
 Update(p *player.Player) error
 Delete(id int) error
}

type TeamRepository interface {
 Save(t *team.Team) error
 FindByID(id int) (*team.Team, error)
 FindAll() ([]*team.Team, error)
 Update(t *team.Team) error
 Delete(id int) error
}

type EventRepository interface {
 Save(e *event.Event) error
 FindAll() ([]event.Event, error)
}

type NewsRepository interface {
 Save(n *news.Article) error
 FindAll() ([]news.Article, error)
}
