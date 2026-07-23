package main

import (
 "fmt"
 "log"

 "github.com/egor-shik/basketball-manager/internal/domain/news"
 "github.com/egor-shik/basketball-manager/internal/domain/player"
 "github.com/egor-shik/basketball-manager/internal/domain/season"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
 "github.com/egor-shik/basketball-manager/internal/repository/sqlite"
 "github.com/egor-shik/basketball-manager/internal/service"
 "github.com/egor-shik/basketball-manager/internal/simulation"
)

func main() {

 db, err := sqlite.New("manager.db")
 if err != nil {
  log.Fatalf("Error connecting to SQLite: %v", err)
 }
 defer db.Close()

 if err := db.InitSchema("database/migrations/001_init.sql"); err != nil {
  log.Fatalf("Migration execution error: %v", err)
 }

 playerRepo := sqlite.NewPlayerRepository(db.DB())
 teamRepo := sqlite.NewTeamRepository(db.DB())
 eventRepo := sqlite.NewEventRepository(db.DB())
 newsRepo := sqlite.NewNewsRepository(db.DB())

 eventService := service.NewEventService()
 newsGen := news.NewGenerator()
 simulator := simulation.NewSimulator()

 p1 := player.NewPlayer(0, "Michael", "Jordan", 23, player.ShootingGuard)
 p1.Contract = player.Contract{Salary: 10_000_000, MarketValue: 50_000_000}
 p2 := player.NewPlayer(0, "LeBron", "James", 28, player.PowerForward)
 p2.Contract = player.Contract{Salary: 12_000_000, MarketValue: 55_000_000}

 _ = playerRepo.Save(p1)
 _ = playerRepo.Save(p2)

 teamA := team.NewTeam("Chicago Foxes", nil, []*player.Player{p1})
 teamA.Budget = team.Budget{Balance: 50_000_000, SalaryCap: 30_000_000, Payroll: 10_000_000}

 teamB := team.NewTeam("Boston Eagles", nil, []*player.Player{p2})
 teamB.Budget = team.Budget{Balance: 40_000_000, SalaryCap: 30_000_000, Payroll: 12_000_000}
 _ = teamRepo.Save(teamA)
 _ = teamRepo.Save(teamB)

 currentSeason := season.NewSeason(1, []*team.Team{teamA, teamB})

 fmt.Println("=== STAGE: SIMULATION AND STORAGE IN SQLITE ===")

 for currentSeason.HasNextMatch() {
  nextMatch := currentSeason.NextMatch()
  result := simulator.PlayMatch(nextMatch.HomeTeam, nextMatch.AwayTeam)
  currentSeason.ApplyResult(nextMatch, result)

  matchEvent := eventService.CreateMatchFinishedEvent(currentSeason.SeasonNumber, nextMatch, result)
  
  if err := eventRepo.Save(&matchEvent); err != nil {
   log.Printf("Error saving the event: %v", err)
  }
 }

 savedEvents, _ := eventRepo.FindAll()
 articles := newsGen.Generate(savedEvents)

 for i := range articles {
  _ = newsRepo.Save(&articles[i])
 }

 fmt.Println("\n=== STEP: READING DATA FROM THE DATABASE ===")

 dbPlayers, _ := playerRepo.FindAll()
 fmt.Printf("Players in the database: %d\n", len(dbPlayers))
 for _, p := range dbPlayers {
  fmt.Printf(" - ID: %d | %s | Position: %s | Salary: %d $\n", p.Personal.ID, p.FullName(), p.Personal.Pos, p.Contract.Salary)
 }

 dbTeams, _ := teamRepo.FindAll()
 fmt.Printf("\nCommands in the database: %d\n", len(dbTeams))
 for _, t := range dbTeams {
  fmt.Printf(" - Team: %s | Balance: %d $\n", t.Name, t.Budget.CurrentBalance())
 }

 dbEvents, _ := eventRepo.FindAll()
 fmt.Printf("\nEvents saved in the database: %d\n", len(dbEvents))

 dbNews, _ := newsRepo.FindAll()
 fmt.Printf("News saved in the database: %d\n", len(dbNews))
 for _, art := range dbNews {
  fmt.Printf("[%s] %s\n", art.Title, art.Body)
 }
}