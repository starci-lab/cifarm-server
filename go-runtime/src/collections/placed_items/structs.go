package collections_placed_items

import (
	collections_seeds "cifarm-server/src/collections/seeds"
)

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type SeedGrowthInfo struct {
	CurrentStage             int                    `json:"currentStage"`
	CurrentStageTimeElapsed  int64                  `json:"currentStageTimeElapsed"`
	TotalTimeElapsed         int64                  `json:"totalTimeElapsed"`
	HarvestQuantityRemaining int                    `json:"harvestQuantityRemaining"`
	IsInfested               bool                   `json:"isInfested"`
	IsWeedy                  bool                   `json:"isWeedy"`
	Seed                     collections_seeds.Seed `json:"seed"`
}

type PlacedItem struct {
	ReferenceId    string         `json:"referenceId"`
	Position       Position       `json:"position"`
	Type           int            `json:"type"`
	SeedGrowthInfo SeedGrowthInfo `json:"seedGrowthInfo"`
	IsPlanted      bool           `json:"isPlanted"`
}
