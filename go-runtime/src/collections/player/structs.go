package collections_player

type TelegramData struct {
	UserId int `json:"userId,omitempty"`
}
type Metadata struct {
	Key            string       `json:"key,omitempty"`
	ChainKey       string       `json:"chainKey,omitempty"`
	AccountAddress string       `json:"accountAddress,omitempty"`
	Network        string       `json:"network,omitempty"`
	TelegramData   TelegramData `json:"telegramData,omitempty"`
}

type VisitState struct {
	Key      string `json:"key,omitempty"`
	UserId   string `json:"userId,omitempty"`
	IsRandom bool   `json:"isRandom,omitempty"`
}

type PlayerStats struct {
	Key          string       `json:"key,omitempty"`
	LevelInfo    LevelInfo    `json:"levelInfo,omitempty"`
	TutorialInfo TutorialInfo `json:"tutorialInfo,omitempty"`
	Invites      []int        `json:"invites,omitempty"`
	EnergyInfo   EnergyInfo   `json:"energyInfo,omitempty"`
}

type TutorialInfo struct {
	TutorialIndex int `json:"tutorialIndex,omitempty"`
	StepIndex     int `json:"stepIndex,omitempty"`
}

type LevelInfo struct {
	Experiences     int `json:"experiences,omitempty"`
	ExperienceQuota int `json:"experienceQuota,omitempty"`
	Level           int `json:"level,omitempty"`
}

type EnergyInfo struct {
	CurrentEnergy     int   `json:"currentEnergy,omitempty"`
	MaxEnergy         int   `json:"maxEnergy,omitempty"`
	EnergyQuota       int   `json:"energyQuota,omitempty"`
	RecoveryTimeCount int64 `json:"recoveryTimeCount,omitempty"`
}

type RewardTracker struct {
	Key              string           `json:"key,omitempty"`
	DailyRewardsInfo DailyRewardsInfo `json:"dailyRewardsInfo,omitempty"`
	SpinInfo         SpinInfo         `json:"spinInfo,omitempty"`
}

type DailyRewardsInfo struct {
	Streak         int   `json:"streak,omitempty"`
	LastClaimTime  int64 `json:"lastClaimTime,omitempty"`
	NumberOfClaims int   `json:"numberOfClaims,omitempty"`
}

type SpinInfo struct {
	LastSpinTime int64 `json:"lastSpinTime,omitempty"`
	SpinCount    int   `json:"spinCount,omitempty"`
}

type Followings struct {
	Key           string                  `json:"key,omitempty"`
	FollowedUsers map[string]FollowedUser `json:"followedUsers,omitempty"`
}

type FollowedUser struct {
	IsPrior bool `json:"isPrior,omitempty"`
}
