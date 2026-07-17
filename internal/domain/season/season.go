package season

import (
 "github.com/egor-shik/basketball-manager/internal/domain/match"
 "github.com/egor-shik/basketball-manager/internal/domain/team"
)

type Season struct {
 SeasonNumber int
 Schedule     *Schedule
 Standings    *Standings
 Finished     bool
}

func NewSeason(seasonNumber int, teams []*team.Team) *Season {
 
 table := make([]Standing, len(teams))
 for i, t := range teams {
  table[i] = Standing{Team: t}
 }

 return &Season{
  SeasonNumber: seasonNumber,
  Schedule:     NewSchedule(teams),
  Standings:    &Standings{Table: table},
  Finished:     false,
 }
}

func (s *Season) HasNextMatch() bool {
 return s.Schedule.HasNext()
}

func (s *Season) NextMatch() *match.Match {
 return s.Schedule.Next()
}

func (s *Season) ApplyResult(m *match.Match, res *match.MatchResult) {
 homeRow := s.Standings.Find(m.HomeTeam)
 awayRow := s.Standings.Find(m.AwayTeam)

 if homeRow == nil || awayRow == nil {
  return 
 }

 homeRow.PointsFor += res.HomeScore
 homeRow.PointsAgainst += res.AwayScore
 awayRow.PointsFor += res.AwayScore
 awayRow.PointsAgainst += res.HomeScore

 if res.Winner == match.WinnerHome {
  homeRow.Wins++
  homeRow.Points += PointsWin

  awayRow.Losses++
  awayRow.Points += PointsLoss
 } else if res.Winner == match.WinnerAway {
  awayRow.Wins++
  awayRow.Points += PointsWin

  homeRow.Losses++
  homeRow.Points += PointsLoss
 }

 s.Standings.Sort()

 if !s.Schedule.HasNext() {
  s.Finished = true
 }
}

func (s *Season) Champion() *team.Team {
 if !s.Finished {
  return nil
 }
 return s.Standings.Leader()
}