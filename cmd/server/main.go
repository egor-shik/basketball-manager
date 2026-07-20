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
 p5 := player.NewPlayer(5, "Steph", "Curry", 25, player.PointGuard)
 p5.SetAllRatings(94)
 p6 := player.NewPlayer(6, "Klay", "Thompson", 27, player.ShootingGuard)
 p6.SetAllRatings(87)
 p7 := player.NewPlayer(7, "Draymond", "Green", 29, player.PowerForward)
 p7.SetAllRatings(85)
 p8 := player.NewPlayer(8, "Devin", "Booker", 25, player.ShootingGuard)
 p8.SetAllRatings(89)
 p9 := player.NewPlayer(9, "Jayson", "Tatum", 26, player.SmallForward)
 p9.SetAllRatings(91)
 p10 := player.NewPlayer(10, "Jaylen", "Brown", 27, player.ShootingGuard)
 p10.SetAllRatings(86)
 p11 := player.NewPlayer(11, "Joel", "Embiid", 28, player.Center)
 p11.SetAllRatings(93)
 p12 := player.NewPlayer(12, "Anthony", "Davis", 30, player.PowerForward)
 p12.SetAllRatings(91)
 p13 := player.NewPlayer(13, "Luka", "Doncic", 24, player.PointGuard)
 p13.SetAllRatings(96)
 p14 := player.NewPlayer(14, "Kyrie", "Irving", 29, player.PointGuard)
 p14.SetAllRatings(90)
 p15 := player.NewPlayer(15, "Jimmy", "Butler", 33, player.SmallForward)
 p15.SetAllRatings(88)

 c1 := coach.NewCoach(1, "Phil", "Jackson", 15)
 c2 := coach.NewCoach(2, "Gregg", "Popovich", 20)
 c3 := coach.NewCoach(3, "Erik", "Spoelstra", 10)
 c4 := coach.NewCoach(4, "Steve", "Kerr", 8)
 c5 := coach.NewCoach(5, "Doc", "Rivers", 17)
 c6 := coach.NewCoach(6, "Mike", "D'Antoni", 14)

 teamA := team.NewTeam("Chicago Foxes", c1, []*player.Player{p1, p2})
 teamA.Budget = team.Budget{Balance: 50_000_000, SalaryCap: 30_000_000, Payroll: 15_000_000}

 teamB := team.NewTeam("Boston Eagles", c2, []*player.Player{p3, p4})
 teamB.Budget = team.Budget{Balance: 12_000_000, SalaryCap: 20_000_000, Payroll: 19_500_000}

 teamC := team.NewTeam("LA Lakers", c3, []*player.Player{p5, p6, p7})
 teamC.Budget = team.Budget{Balance: 45_000_000, SalaryCap: 35_000_000, Payroll: 22_000_000}

 teamD := team.NewTeam("Golden State Warriors", c4, []*player.Player{p8, p9, p10})
 teamD.Budget = team.Budget{Balance: 35_000_000, SalaryCap: 30_000_000, Payroll: 18_000_000}

 teamE := team.NewTeam("Philadelphia 76ers", c5, []*player.Player{p11, p12})
 teamE.Budget = team.Budget{Balance: 28_000_000, SalaryCap: 25_000_000, Payroll: 16_000_000}

 teamF := team.NewTeam("Dallas Mavericks", c6, []*player.Player{p13, p14, p15})
 teamF.Budget = team.Budget{Balance: 32_000_000, SalaryCap: 28_000_000, Payroll: 20_000_000}

 teams := []*team.Team{teamA, teamB, teamC, teamD, teamE, teamF}

 freeAgent1 := player.NewPlayer(16, "Giannis", "Antetokounmpo", 27, player.PowerForward)
 freeAgent1.Contract = player.Contract{Salary: 8_000_000, YearsLeft: 2}
 freeAgent1.SetAllRatings(97)

 freeAgent2 := player.NewPlayer(17, "Nikola", "Jokic", 28, player.Center)
 freeAgent2.Contract = player.Contract{Salary: 6_000_000, YearsLeft: 1}
 freeAgent2.SetAllRatings(95)

 freeAgent3 := player.NewPlayer(18, "Kawhi", "Leonard", 31, player.SmallForward)
 freeAgent3.Contract = player.Contract{Salary: 4_000_000, YearsLeft: 3}
 freeAgent3.SetAllRatings(92)

 fmt.Printf("[Transfer] Signing attempt %s to %s...\n", freeAgent1.FullName(), teamA.Name)
 if err := transferService.SignPlayer(teamA, freeAgent1); err == nil {
  ev := eventService.CreatePlayerSignedEvent(teamA.Name, freeAgent1.FullName())
  eventStore = append(eventStore, ev)
  fmt.Println("Successfully signed in!")
 }

 fmt.Printf("[Transfer] Signing attempt %s to %s...\n", freeAgent2.FullName(), teamC.Name)
 if err := transferService.SignPlayer(teamC, freeAgent2); err == nil {
  ev := eventService.CreatePlayerSignedEvent(teamC.Name, freeAgent2.FullName())
  eventStore = append(eventStore, ev)
  fmt.Println("Successfully signed in!")
 }

 fmt.Printf("[Transfer] Signing attempt %s to %s...\n", freeAgent3.FullName(), teamE.Name)
 if err := transferService.SignPlayer(teamE, freeAgent3); err == nil {
  ev := eventService.CreatePlayerSignedEvent(teamE.Name, freeAgent3.FullName())
  eventStore = append(eventStore, ev)
  fmt.Println("Successfully signed in!")
 }

 fmt.Printf("[Transfer] Signing attempt %s to %s (already signed)...\n", freeAgent1.FullName(), teamB.Name)
 if err := transferService.SignPlayer(teamB, freeAgent1); err == nil {
  ev := eventService.CreatePlayerSignedEvent(teamB.Name, freeAgent1.FullName())
  eventStore = append(eventStore, ev)
  fmt.Println("Successfully signed in!")
 } else {
  fmt.Printf("Failed: %v\n", err)
 }

 fmt.Printf("[Transfer] Signing attempt %s to %s...\n", freeAgent2.FullName(), teamD.Name)
 if err := transferService.SignPlayer(teamD, freeAgent2); err == nil {
  ev := eventService.CreatePlayerSignedEvent(teamD.Name, freeAgent2.FullName())
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