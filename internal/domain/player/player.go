package player

type Player struct {
	Personal  PersonalInfo
	Ratings   Ratings
	Potential Potential
	Stats     Statistics
	Mood      Mood
	Contract  Contract
}

type PersonalInfo struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	HeightCm  int
	WeightKg  int
	Pos       Position
}

type Ratings struct {
	ThreePoint   float64 // 0-100
	MidRange     float64 // 0-100
	Inside       float64 // 0-100
	Passing      float64 // 0-100
	Defense      float64 // 0-100
	Rebounding   float64 // 0-100
	Athleticism  float64 // 0-100
	BasketballIQ float64 // 0-100
}

type Potential struct {
	PotentialOverall float64
}

type Mood struct {
	Value int // 0-100
	Ego   int // 0-5
}

type Position string

const (
	PointGuard    Position = "PG"
	ShootingGuard Position = "SG"
	SmallForward  Position = "SF"
	PowerForward  Position = "PF"
	Center        Position = "C"
)

func (p *Player) FullName() string {
	return p.Personal.FirstName + " " + p.Personal.LastName
}

func (p *Player) Overall() float64 {
	// Calculate all rating of player
	r := p.Ratings
	return (r.ThreePoint + r.MidRange + r.Inside +
		r.Passing + r.Defense + r.Rebounding +
		r.Athleticism + r.BasketballIQ) / 8
}

func (p *Player) GetRating() Ratings {
	return p.Ratings
}

func (p *Player) SetThreePoint(value float64) {
	p.Ratings.ThreePoint = clamp(value, 0, 100)
}

func (p *Player) SetAllRatings(value float64) {
	v := clamp(value, 0, 100)
	p.Ratings.ThreePoint = v
	p.Ratings.MidRange = v
	p.Ratings.Inside = v
	p.Ratings.Passing = v
	p.Ratings.Defense = v
	p.Ratings.Rebounding = v
	p.Ratings.Athleticism = v
	p.Ratings.BasketballIQ = v
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
