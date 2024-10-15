package collections_daily_rewards

//amount on indivial day
type DailyReward struct {
	Key       string `json:"key"`
	Amount    int64  `json:"amount"`
	Day       int    `json:"day"`
	IsLastDay bool   `json:"isLastDay"`
	//can be nil if isLastDay = false
	DailyRewardPossibilities map[int]LastDailyRewardPossibility `json:"dailyRewardPossibilities"`
}

//random reward
//either gold
//small chance to get the CARROT tokens!
type LastDailyRewardPossibility struct {
	GoldAmount  int64   `json:"goldAmount"`
	TokenAmount float64 `json:"tokenAmount"`
	//to indecate if random 100, the range threshold acheive the reward
	ThresholdMin float64 `json:"thresholdMin"`
	ThresholdMax float64 `json:"thresholdMax"`
}
