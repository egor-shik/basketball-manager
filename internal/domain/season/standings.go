package season

import (
 "sort"

 "github.com/egor-shik/basketball-manager/internal/domain/team"
)

const (
 PointsWin  = 2
 PointsLoss = 1
)

type Standing struct {
 Team   *team.Team
 Wins   int
 Losses int
 Points int

 PointsFor     int 
 PointsAgainst int 
}

type Standings struct {
 Table []Standing
}

func (st *Standings) Find(t *team.Team) *Standing {
 for i := range st.Table {
  if st.Table[i].Team == t {
   return &st.Table[i]
  }
 }
 return nil
}

func (st *Standings) Sort() {
 sort.Slice(st.Table, func(i, j int) bool {
  if st.Table[i].Points != st.Table[j].Points {
   return st.Table[i].Points > st.Table[j].Points
  }
  return st.Table[i].Wins > st.Table[j].Wins
 })
}

func (st *Standings) Leader() *team.Team {
 if len(st.Table) == 0 {
  return nil
 }
 return st.Table[0].Team
}