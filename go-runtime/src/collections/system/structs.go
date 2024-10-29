package collections_system

type Users struct {
	Key     string   `json:"key,omitempty"`
	UserIds []string `json:"userIds,omitempty"`
}

type LastServerUptime struct {
	Key           string `json:"key,omitempty"`
	TimeInSeconds int64  `json:"timeInSeconds,omitempty"`
}

type MatchInfo struct {
	Key            string `json:"key,omitempty"`
	CentralMatchId string `json:"centralMatchId,omitempty"`
	TimerMatchId   string `json:"timerMatchId,omitempty"`
}

type SpeedUp struct {
	Key                 string `json:"key,omitempty"`
	AnimalProcedureTime int64  `json:"animalProcedureTime,omitempty"`
	SeedGrowthTime      int64  `json:"time,omitempty"`
	EnergyGain          int64  `json:"energyGain,omitempty"`
}

type ActivityInfo struct {
	ExperiencesGain int `json:"experiencesGain,omitempty"`
	EnergyCost      int `json:"energyCost,omitempty"`
}
type Activities struct {
	Key                string       `json:"key,omitempty"`
	Water              ActivityInfo `json:"water,omitempty"`
	FeedAnimal         ActivityInfo `json:"feedAnimal,omitempty"`
	UsePestiside       ActivityInfo `json:"usePestiside,omitempty"`
	UseFertilizer      ActivityInfo `json:"useFertilizer,omitempty"`
	UseHerbicide       ActivityInfo `json:"useHerbicide,omitempty"`
	HelpUseHerbicide   ActivityInfo `json:"helpUseHerbicide,omitempty"`
	HelpUsePestiside   ActivityInfo `json:"helpUsePestiside,omitempty"`
	HelpWater          ActivityInfo `json:"helpWater,omitempty"`
	ThiefCrop          ActivityInfo `json:"thiefCrop,omitempty"`
	ThiefAnimalProduct ActivityInfo `json:"thiefAnimalProduct,omitempty"`
	CureAnimal         ActivityInfo `json:"cureAnimal,omitempty"`
	HelpCureAnimal     ActivityInfo `json:"helpCureAnimal,omitempty"`
}

type Rewards struct {
	Key         string      `json:"key,omitempty"`
	FromInvites FromInvites `json:"fromInvites,omitempty"`
	Referred    int64       `json:"referred,omitempty"`
}

type FromInvites struct {
	Key     string         `json:"key,omitempty"`
	Metrics map[int]Metric `json:"metrics,omitempty"`
}

type Metric struct {
	Key   int   `json:"key,omitempty"`
	Value int64 `json:"value,omitempty"`
}

type CropRandomness struct {
	Key               string  `json:"key,omitempty"`
	Theif3            float64 `json:"theif3,omitempty"`
	Theif2            float64 `json:"theif2,omitempty"`
	NeedWater         float64 `json:"needWater,omitempty"`
	IsWeedyOrInfested float64 `json:"isWeedyOrInfested,omitempty"`
}

type TokenConfigure struct {
	Key string `json:"key,omitempty"`
	//token in-game decimal
	Decimals int `json:"decimals,omitempty"`
}

type StarterConfigure struct {
	Key        string `json:"key,omitempty"`
	GoldAmount int64  `json:"goldAmount,omitempty"`
}

type SpinConfigure struct {
	Key          string `json:"key,omitempty"`
	SpinPrice    int64  `json:"spinPrice,omitempty"`
	FreeSpinTime int64  `json:"freeSpinTime,omitempty"`
}
