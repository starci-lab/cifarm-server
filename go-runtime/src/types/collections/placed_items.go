package collections

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type SeedGrowthInfo struct {
	CurrentStage             int       `json:"currentStage"`
	CurrentStageTimeElapsed  float32   `json:"currentStageTimeElapsed"`
	TotalTimeElapsed         float32   `json:"totalTimeElapsed"`
	HarvestQuantityRemaining int       `json:"harvestQuantityRemaining"`
	IsInfested               bool      `json:"isInfested"`
	IsWeedy                  bool      `json:"isWeedy"`
	PlantSeed                PlantSeed `json:"plantSeed"`
}

type PlacedItem struct {
	Id             string         `json:"id"`
	Position       Position       `json:"position"`
	Type           int            `json:"type"`
	SeedGrowthInfo SeedGrowthInfo `json:"seedGrowthInfo"`
}
