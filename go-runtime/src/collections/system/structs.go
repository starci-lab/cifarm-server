package collections_system

type Users struct {
	UserIds []string `json:"userIds"`
}

type LastServerUptime struct {
	TimeInSeconds int64 `json:"timeInSeconds"`
}

type CentralMatchInfo struct {
	MatchId string `json:"matchId"`
}

type SpeedUp struct {
	Time int64 `json:"time"`
}
