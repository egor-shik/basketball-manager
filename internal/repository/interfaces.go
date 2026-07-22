package repository 

type PlayerRepository interface {
    Save(*player.Player) error
    FindByID(id int) (*player.Player, error)
    FindAll() ([]*player.Player, error)
    Update(*player.Player) error
}

type TeamRepository interface {
    Save(*player.Player) error
    FindByID(id int) (*player.Player, error)
    FindAll() ([]*player.Player, error)
    Update(*player.Player) error
}

type SeasonRepository interface {
    Save(*player.Player) error
    FindByID(id int) (*player.Player, error)
    FindAll() ([]*player.Player, error)
    Update(*player.Player) error
}

type EventRepository interface {
    Save(*player.Player) error
    FindByID(id int) (*player.Player, error)
    FindAll() ([]*player.Player, error)
    Update(*player.Player) error
}

type NewsRepository interface {
    Save(*player.Player) error
    FindByID(id int) (*player.Player, error)
    FindAll() ([]*player.Player, error)
    Update(*player.Player) error
}