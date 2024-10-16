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
	EnergyGain          int64  `json:"energyGain"`
}

type ActivityInfo struct {
	ExperiencesGain int `json:"experiencesGain"`
	EnergyCost      int `json:"energyCost"`
}
type Activities struct {
	Key                string       `json:"key"`
	Water              ActivityInfo `json:"water"`
	FeedAnimal         ActivityInfo `json:"feedAnimal"`
	UsePestiside       ActivityInfo `json:"usePestiside"`
	UseFertilizer      ActivityInfo `json:"useFertilizer"`
	UseHerbicide       ActivityInfo `json:"useHerbicide"`
	HelpUseHerbicide   ActivityInfo `json:"helpUseHerbicide"`
	HelpUsePestiside   ActivityInfo `json:"helpUsePestiside"`
	HelpWater          ActivityInfo `json:"helpWater"`
	ThiefCrop          ActivityInfo `json:"thiefCrop"`
	ThiefAnimalProduct ActivityInfo `json:"thiefAnimalProduct"`
}

type Rewards struct {
	Key         string      `json:"key"`
	FromInvites FromInvites `json:"fromInvites"`
	Referred    int64       `json:"referred"`
}

type FromInvites struct {
	Key     string         `json:"key"`
	Metrics map[int]Metric `json:"metrics"`
}

type Metric struct {
	Key   int   `json:"key"`
	Value int64 `json:"value"`
}

type CropRandomness struct {
	Key               string  `json:"key"`
	Theif3            float64 `json:"theif3"`
	Theif2            float64 `json:"theif2"`
	NeedWater         float64 `json:"needWater"`
	IsWeedyOrInfested float64 `json:"isWeedyOrInfested"`
}

type TokenConfigure struct {
	Key string `json:"key"`
	//token in-game decimal
	Decimals int `json:"decimals"`
}

type StarterConfigure struct {
	Key        string `json:"key"`
	GoldAmount int64  `json:"goldAmount"`
}

type SpinConfigure struct {
	Key          string `json:"key"`
	SpinPrice    int64  `json:"spinPrice"`
	FreeSpinTime int64  `json:"freeSpinTime"`
}
