package main

import (
 "fmt"

 "github.com/egor-shik/basketball-manager/internal/domain/coach"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
 "github.com/egor-shik/basketball-manager/internal/simulation"
)

func setSkills(p *player.Player, three, mid, inside, pass, def, reb, ath, iq float64) {
 p.SetThreePoint(three)
}

func main() {
 // team A
p1 := player.NewPlayer(1, "John", "Smith", 25, player.PointGuard)
p1.SetAllRatings(95) 

p2 := player.NewPlayer(2, "Michael", "Jordan", 28, player.ShootingGuard)
p2.SetAllRatings(98) 

coachA := coach.NewCoach(1, "Phil", "Jackson", 20)
coachA.Buffs.ShootingBonus = 0.1 

teamA := team.NewTeam("Chicago Foxes", coachA, []*player.Player{p1, p2})
teamA.Morale.Chemistry = 1.0 


// team B
p3 := player.NewPlayer(3, "Kevin", "Durant", 27, player.SmallForward)
p3.SetAllRatings(10) 

p4 := player.NewPlayer(4, "LeBron", "James", 32, player.PowerForward)
p4.SetAllRatings(10) 

coachB := coach.NewCoach(2, "Gregg", "Popovich", 25) 

teamB := team.NewTeam("Boston Eagles", coachB, []*player.Player{p3, p4})
teamB.Morale.Chemistry = 0.1 

 simulator := simulation.NewSimulator()
 result := simulator.PlayMatch(teamA, teamB)

 fmt.Printf("%s %d\n", teamA.Name, result.HomeScore)
 fmt.Printf("%s %d\n", teamB.Name, result.AwayScore)
 fmt.Println("--------------------------------")
 
 allPlayers := map[int]string{
  1: p1.FullName(), 2: p2.FullName(),
  3: p3.FullName(), 4: p4.FullName(),
 }

 fmt.Printf("MVP: %s\n", allPlayers[result.MVPID])
 fmt.Println("--------------------------------")

 for _, stat := range result.PlayerStats {
  fmt.Printf("%s\n%d PTS | %d AST | %d REB\n", allPlayers[stat.PlayerID], stat.Points, stat.Assists, stat.Rebounds)
  fmt.Println("--------------------------------")
 }
}