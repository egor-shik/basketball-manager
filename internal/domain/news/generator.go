package news

import (
 "fmt"

 "github.com/egor-shik/basketball-manager/internal/domain/event"
)

type Generator struct {
 currentID int
}

func NewGenerator() *Generator {
 return &Generator{currentID: 1}
}

func (g *Generator) nextID() int {
 id := g.currentID
 g.currentID++
 return id
}

func (g *Generator) Generate(events []event.Event) []Article {
 var articles []Article
 for _, e := range events {
  if article, ok := g.generateSingle(e); ok {
   articles = append(articles, article)
  }
 }
 return articles
}

func (g *Generator) generateSingle(e event.Event) (Article, bool) {
 var title, body string
//maybe convert it to a map later 
 switch e.Type {
 case event.MatchFinished:
  title = "MATCH RESULTS"
  if e.HomeScore > e.AwayScore {
   body = fmt.Sprintf("Club %s won a convincing home victory over %s with a score of %d:%d!", e.TeamName, e.OpponentName, e.HomeScore, e.AwayScore)
  } else {
   body = fmt.Sprintf("%s suffers an embarrassing defeat in their home arena. %s takes home the victory with a score of %d:%d.", e.TeamName, e.OpponentName, e.AwayScore, e.HomeScore)
  }

 case event.PlayerSigned:
  title = "SIGNING!!!"
  body = fmt.Sprintf("Official: %s have strengthened their squad! The club has signed a contract with player %s.", e.TeamName, e.PlayerName)

 case event.PlayerReleased:
  title = "CHANGES IN THE ROSTER"
  body = fmt.Sprintf("%s officially announces the termination of the contract with free agent %s.", e.TeamName, e.PlayerName)

 case event.SeasonFinished:
  title = "THE CHAMPION IS DETERMINED!"
  body = fmt.Sprintf("Season %d has come to an end. %s is the champion of the league and takes home the gold cup!", e.SeasonNumber, e.TeamName)

 default:
  return Article{}, false
 }
 return Article{
	ID:             g.nextID(),
	Title:          title,
	Body:           body,
	CreatedAt:      e.Timestamp,
	RelatedEventID: e.ID,
   }, true
  }