package match

type WinnerSide int

const (
	WinnerNone WinnerSide = iota
	WinnerHome
	WinnerAway
)

type MatchResult struct {
	HomeScore   int
	AwayScore   int
	Winner      WinnerSide
	MVPID       int 
	PlayerStats []PlayerMatchStats
	Finished    bool
}
