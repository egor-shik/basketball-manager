package player


type Statistics struct {
    Games int 
    Points int
    Assists int
    Rebounds int 
}

func (p *Player) PointsPerGame() float64 {
    if p.Stats.Games == 0 {
        return 0
    }
    return  float64(p.Stats.Points) / float64(p.Stats.Games)
}

func (p *Player) AssistsPerGame() float64 {
    if p.Stats.Games == 0 {
        return 0
    }
    return float64(p.Stats.Assists) / float64(p.Stats.Games)
}

func (p *Player) ReboundsPerGame() float64 {
    if p.Stats.Games == 0 {
        return 0
    }
    return float64(p.Stats.Rebounds) / float64(p.Stats.Games)
}