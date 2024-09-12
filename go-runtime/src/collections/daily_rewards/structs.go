package collections_daily_rewards

type DailyReward struct {
	Key    string `json:"key"`
	Amount int64  `json:"amount"`
	Days   int    `json:"days"`
}
