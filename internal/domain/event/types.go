package event

type Type string

const (
	MatchFinished  Type = "MATCH_FINISHED"
	PlayerSigned   Type = "PLAYER_SIGNED"
	PlayerReleased Type = "PLAYER_RELEASED"
	TradeCompleted Type = "TRADE_COMPLETED"
	SeasonFinished Type = "SEASON_FINISHED"
)
