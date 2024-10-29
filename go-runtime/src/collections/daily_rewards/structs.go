package collections_daily_rewards

//amount on indivial day
type DailyReward struct {
	Key       string `json:"key,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
	Day       int    `json:"day,omitempty"`
	IsLastDay bool   `json:"isLastDay,omitempty"`
	//can be nil if isLastDay = false
	DailyRewardPossibilities map[int]LastDailyRewardPossibility `json:"dailyRewardPossibilities,omitempty"`
}

//random reward
//either gold
//small chance to get the CARROT tokens!
type LastDailyRewardPossibility struct {
	Key         string  `json:"key,omitempty"`
	GoldAmount  int64   `json:"goldAmount,omitempty"`
	TokenAmount float64 `json:"tokenAmount,omitempty"`
	//to indecate if random 100, the range threshold acheive the reward
	ThresholdMin float64 `json:"thresholdMin,omitempty"`
	ThresholdMax float64 `json:"thresholdMax,omitempty"`
}
