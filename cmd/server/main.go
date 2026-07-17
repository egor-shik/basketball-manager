package main

import (
 "fmt"

 "github.com/egor-shik/basketball-manager/internal/domain/coach"
 "github.com/egor-shik/basketball-manager/internal/domain/event"
 "github.com/egor-shik/basketball-manager/internal/domain/news"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
 "github.com/egor-shik/basketball-manager/internal/domain/season"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
 "github.com/egor-shik/basketball-manager/internal/service"
 "github.com/egor-shik/basketball-manager/internal/simulation"
)

func main() {

 transferService := service.NewTransferService()
 eventService := service.NewEventService()
 newsGen := news.NewGenerator()

 var eventStore []event.Event

 p1 := player.NewPlayer(1, "John", "Smith", 24, player.PointGuard)
 p1.SetAllRatings(80)
 p2 := player.NewPlayer(2, "Michael", "Jordan", 23, player.ShootingGuard)
 p2.SetAllRatings(95)
 p3 := player.NewPlayer(3, "Kevin", "Durant", 26, player.SmallForward)
 p3.SetAllRatings(90)
 p4 := player.NewPlayer(4, "LeBron", "James", 28, player.PowerForward)
 p4.SetAllRatings(92)

 c1 := coach.NewCoach(1, "Phil", "Jackson", 15)
 teamA := team.NewTeam("Chicago Foxes", c1, []*player.Player{p1, p2})
 teamA.Budget = team.Budget{Balance: 50_000_000, SalaryCap: 30_000_000, Payroll: 15_000_000}

 teamB := team.NewTeam("Boston Eagles", nil, []*player.Player{p3, p4})
 teamB.Budget = team.Budget{Balance: 12_000_000, SalaryCap: 20_000_000, Payroll: 19_500_000}

 teams := []*team.Team{teamA, teamB}

 freeAgent := player.NewPlayer(5, "Steph", "Curry", 25, player.PointGuard)
 freeAgent.Contract = player.Contract{Salary: 5_000_000, YearsLeft: 2}

 fmt.Printf("\n[Transfer] Signing attempt %s...\n", freeAgent.FullName())
 if err := transferService.SignPlayer(teamA, freeAgent); err == nil {

  ev := eventService.CreatePlayerSignedEvent(teamA.Name, freeAgent.FullName())
  eventStore = append(eventStore, ev)
  fmt.Println("Successfully signed in!")
 }

 currentSeason := season.NewSeason(1, teams)
 simulator := simulation.NewSimulator()

 fmt.Println("\n[Season] The start of the championship games...")
 for currentSeason.HasNextMatch() {
  nextMatch := currentSeason.NextMatch()

  result := simulator.PlayMatch(nextMatch.HomeTeam, nextMatch.AwayTeam)
  currentSeason.ApplyResult(nextMatch, result)

  ev := eventService.CreateMatchFinishedEvent(currentSeason.SeasonNumber, nextMatch, result)
  eventStore = append(eventStore, ev)
 }

 if currentSeason.Finished {
  champion := currentSeason.Champion()
  ev := eventService.CreateSeasonFinishedEvent(currentSeason.SeasonNumber, champion.Name)
  eventStore = append(eventStore, ev)
 }

 fmt.Println("\n=== GENERATED EVENTS (EVENT STORE) ===")
 for _, ev := range eventStore {
  fmt.Printf("[%s] ID: %d\n", ev.Type, ev.ID)
 }

 articles := newsGen.Generate(eventStore)

 fmt.Println("\n=== NEWS FEED(GENERATED NEWS) ===")
 for _, art := range articles {
  fmt.Printf("[%s] \n   %s\n   (Published: %s | Based on Event ID: %d)\n\n",
   art.Title, art.Body, art.CreatedAt.Format("15:04:05"), art.RelatedEventID,
  )
 }
}