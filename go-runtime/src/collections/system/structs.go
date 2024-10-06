package collections_system

type Users struct {
	Key     string   `json:"key"`
	UserIds []string `json:"userIds"`
}

type LastServerUptime struct {
	Key           string `json:"key"`
	TimeInSeconds int64  `json:"timeInSeconds"`
}

type MatchInfo struct {
	Key            string `json:"key"`
	CentralMatchId string `json:"centralMatchId"`
	TimerMatchId   string `json:"timerMatchId"`
}

type SpeedUp struct {
	Key                 string `json:"key"`
	AnimalProcedureTime int64  `json:"animalProcedureTime"`
	SeedGrowthTime      int64  `json:"time"`
}

type ActivityExperiences struct {
	Key                string `json:"key"`
	Water              int    `json:"water"`
	UsePestiside       int    `json:"usePestiside"`
	UseFertilizer      int    `json:"useFertilizer"`
	UseHerbicide       int    `json:"useHerbicide"`
	HelpUseHerbicide   int    `json:"helpUseHerbicide"`
	HelpUsePestiside   int    `json:"helpUsePestiside"`
	HelpWater          int    `json:"helpWater"`
	ThiefCrop          int    `json:"thiefCrop"`
	HelpFeedAnimal     int    `json:"helpFeedAnimal"`
	ThiefAnimalProduct int    `json:"thiefAnimalProduct"`
}

type Rewards struct {
	Key         string      `json:"key"`
	FromInvites FromInvites `json:"fromInvites"`
}

type FromInvites struct {
	Key     string         `json:"key"`
	Metrics map[int]Metric `json:"metrics"`
}

type Metric struct {
	Key   int   `json:"key"`
	Value int64 `json:"value"`
}
