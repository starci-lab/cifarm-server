package collections_system

type Users struct {
	Key     string   `json:"key"`
	UserIds []string `json:"userIds"`
}

type LastServerUptime struct {
	Key           string `json:"key"`
	TimeInSeconds int64  `json:"timeInSeconds"`
}

type CentralMatchInfo struct {
	Key     string `json:"key"`
	MatchId string `json:"matchId"`
}

type SpeedUp struct {
	Key  string `json:"key"`
	Time int64  `json:"time"`
}
