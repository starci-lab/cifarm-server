package collections_config

type TelegramData struct {
	UserId int `json:"userId"`
}
type Metadata struct {
	Key            string       `json:"key"`
	ChainKey       string       `json:"chainKey"`
	AccountAddress string       `json:"accountAddress"`
	Network        string       `json:"network"`
	TelegramData   TelegramData `json:"telegramData"`
}

type VisitState struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}

type PlayerStats struct {
	Key          string       `json:"key"`
	LevelInfo    LevelInfo    `json:"levelInfo"`
	TutorialInfo TutorialInfo `json:"tutorialInfo"`
	Invites      []int        `json:"invites"`
	EnergyInfo   EnergyInfo   `json:"energyInfo"`
}

type TutorialInfo struct {
	TutorialIndex int `json:"tutorialIndex"`
	StepIndex     int `json:"stepIndex"`
}

type LevelInfo struct {
	Experiences     int `json:"experiences"`
	ExperienceQuota int `json:"experienceQuota"`
	Level           int `json:"level"`
}

type EnergyInfo struct {
	CurrentEnergy     int   `json:"currentEnergy"`
	MaxEnergy         int   `json:"maxEnergy"`
	EnergyQuota       int   `json:"energyQuota"`
	RecoveryTimeCount int64 `json:"recoveryTimeCount"`
}
