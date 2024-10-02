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
	Key  string `json:"key"`
	Time int64  `json:"time"`
}

type ActivityExperiences struct {
	Water            int64 `json:"water"`
	UsePestiside     int64 `json:"usePestiside"`
	UseFertilizer    int64 `json:"useFertilizer"`
	UseHerbicide     int64 `json:"useHerbicide"`
	HelpUseHerbicide int64 `json:"helpUseHerbicide"`
	HelpUsePestiside int64 `json:"helpUsePestiside"`
	HelpWater        int64 `json:"helpWater"`
	ThiefCrop        int64 `json:"thiefCrop"`
}
