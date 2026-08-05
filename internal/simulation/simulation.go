package simulation

import (
	"math/rand"
	"time"

	"github.com/egor-shik/basketball-manager/internal/domain/match"
	"github.com/egor-shik/basketball-manager/internal/domain/team"
)

// Simulator encapsulates the algorithmic logic for simulating basketball matches and player performances
type Simulator struct {
	rnd *rand.Rand
}

// NewSimulator initializes a Simulator with a unique pseudo-random number generator source
func NewSimulator() *Simulator {
	return &Simulator{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// PlayMatch calculates team strengths, generates final scores, determines the winner
// computes individual player statistics, and identifies the MVP
func (s *Simulator) PlayMatch(homeTeam, awayTeam *team.Team) *match.MatchResult {
	homeRating := s.calculateTeamStrength(homeTeam)
	awayRating := s.calculateTeamStrength(awayTeam)

	homeScore, awayScore := s.generateScore(homeRating, awayRating)

	winner := match.WinnerHome
	if awayScore > homeScore {
		winner = match.WinnerAway
	}

	playerStats := s.generatePlayerStats(homeTeam, homeScore, awayTeam, awayScore)

	mvpID := 0
	maxPoints := -1
	for _, stat := range playerStats {
		if stat.Points > maxPoints {
			maxPoints = stat.Points
			mvpID = stat.PlayerID
		}
	}

	return &match.MatchResult{
		HomeScore:   homeScore,
		AwayScore:   awayScore,
		Winner:      winner,
		MVPID:       mvpID,
		PlayerStats: playerStats,
		Finished:    true,
	}
}

// calculateTeamStrength computes the aggregated strength rating of a team
// incorporating player overalls, coach tactical bonuses, and team chemistry
func (s *Simulator) calculateTeamStrength(t *team.Team) float64 {
	if len(t.Players) == 0 {
		return 0.0
	}

	var totalOverall float64
	for _, p := range t.Players {
		pOverall := p.Overall()

		if t.Coach != nil {
			buffs := t.Coach.Buffs
			coachBonus := (buffs.ShootingBonus + buffs.PassingBonus + buffs.DefenseBonus) * 10
			pOverall += coachBonus
		}
		totalOverall += pOverall
	}

	avgOverall := totalOverall / float64(len(t.Players))
	chemistryBonus := t.Morale.Chemistry * 5.0

	return avgOverall + chemistryBonus
}

// generateScore calculates final match scores based on team rating differences
// random noise, home-field advantage, and resolves potential overtimes
func (s *Simulator) generateScore(homeRating, awayRating float64) (int, int) {
	baseHomeScore := 100.0
	baseAwayScore := 100.0

	ratingDiff := homeRating - awayRating
	baseHomeScore += ratingDiff * 0.5
	baseAwayScore -= ratingDiff * 0.5

	homeNoise := (s.rnd.Float64() * 24) - 12
	awayNoise := (s.rnd.Float64() * 24) - 12
	homeFieldAdvantage := 3.0

	finalHome := int(baseHomeScore + homeNoise + homeFieldAdvantage)
	finalAway := int(baseAwayScore + awayNoise)

	// Resolve tie scores via overtime simulation
	if finalHome == finalAway {
		if s.rnd.Intn(2) == 0 {
			finalHome++
		} else {
			finalAway++
		}
	}

	// Enforce minimum score boundaries to prevent unrealistic low-scoring outcomes
	if finalHome < 60 {
		finalHome = 60
	}
	if finalAway < 60 {
		finalAway = 60
	}

	return finalHome, finalAway
}

// generatePlayerStats distributes match points, assists, and rebounds among team players 
// proportionally based on individual ratings and randomized performance luck
func (s *Simulator) generatePlayerStats(home *team.Team, homeScore int, away *team.Team, awayScore int) []match.PlayerMatchStats {
	var allStats []match.PlayerMatchStats

	distribute := func(t *team.Team, totalScore int) []match.PlayerMatchStats {
		tStats := make([]match.PlayerMatchStats, 0, len(t.Players))
		if len(t.Players) == 0 {
			return tStats
		}

		totalOverall := 0.0
		for _, p := range t.Players {
			totalOverall += p.Overall()
		}

		pointsLeft := totalScore
		for i, p := range t.Players {
			if i == len(t.Players)-1 {
				tStats = append(tStats, match.PlayerMatchStats{
					PlayerID: p.Personal.ID,
					Points:   pointsLeft,
					Assists:  s.rnd.Intn(6) + int(p.Ratings.Passing/20),
					Rebounds: s.rnd.Intn(8) + int(p.Ratings.Rebounding/15),
				})
				break
			}

			pOverall := p.Overall()
			currentTotal := totalOverall

			if currentTotal == 0 {
				pOverall = 50.0
				currentTotal = float64(len(t.Players)) * 50.0
			}
			share := pOverall / currentTotal
			playerLuck := 0.8 + (s.rnd.Float64() * 0.2)

			pPoints := int(float64(totalScore) * share * playerLuck)
			if pPoints > pointsLeft {
				pPoints = pointsLeft
			}
			pointsLeft -= pPoints
			if pointsLeft < 0 {
				pointsLeft = 0
			}

			tStats = append(tStats, match.PlayerMatchStats{
				PlayerID: p.Personal.ID,
				Points:   pPoints,
				Assists:  s.rnd.Intn(5) + int(p.Ratings.Passing/25),
				Rebounds: s.rnd.Intn(7) + int(p.Ratings.Rebounding/18),
			})
		}
		return tStats
	}

	allStats = append(allStats, distribute(home, homeScore)...)
	allStats = append(allStats, distribute(away, awayScore)...)

	return allStats
}
