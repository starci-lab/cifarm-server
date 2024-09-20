package collections_placed_items

import (
	collections_crops "cifarm-server/src/collections/crops"
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
	Crop                     collections_crops.Crop `json:"crop"`
	PlantCurrentState        int                    `json:"plantCurrentState"`
	ThiefedBy                []string               `json:"thiefedBy"`
}

type PlacedItem struct {
	Key            string         `json:"key"`
	ReferenceKey   string         `json:"referenceKey"`
	InventoryKey   string         `json:"inventoryKey"`
	Position       Position       `json:"position"`
	Type           int            `json:"type"`
	SeedGrowthInfo SeedGrowthInfo `json:"seedGrowthInfo"`
	IsPlanted      bool           `json:"isPlanted"`
	FullyMatured   bool           `json:"fullyMatured"`
}
