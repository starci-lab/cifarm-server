package collections_config

type Metadata struct {
	Key            string `json:"key"`
	ChainKey       string `json:"chainKey"`
	AccountAddress string `json:"accountAddress"`
	Network        string `json:"network"`
}

type VisitState struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}

type PlayerStats struct {
	Key             string `json:"key"`
	Experiences     int    `json:"experiences"`
	ExperienceQuota int    `json:"experienceQuota"`
	Level           int    `json:"level"`
	TutorialIndex   int    `json:"tutorialIndex"`
	StepIndex       int    `json:"stepIndex"`
}
