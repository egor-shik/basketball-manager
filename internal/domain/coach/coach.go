package coach

import "github.com/egor-shik/basketball-manager/internal/domain/player"

type Coach struct {
	Id         int
	FirstName  string
	LastName   string
	Experience int // years
	Buffs      CoachBuffs
	Contract   player.Contract
}

type CoachBuffs struct {
	DefenseBonus  float64 // 0.0 - 0.2 (20% bonus)
	PassingBonus  float64
	ShootingBonus float64
	MoraleBonus   float64
}

func (c *Coach) FullName() string {
	return c.FirstName + " " + c.LastName
}
