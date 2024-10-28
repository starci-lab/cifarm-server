package collections_placed_items

import (
	collections_animals "cifarm-server/src/collections/animals"
	collections_buildings "cifarm-server/src/collections/buildings"
	collections_crops "cifarm-server/src/collections/crops"
)

type Position struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
}

type SeedGrowthInfo struct {
	CurrentStage             int                    `json:"currentStage,omitempty"`
	CurrentStageTimeElapsed  int64                  `json:"currentStageTimeElapsed,omitempty"`
	TotalTimeElapsed         int64                  `json:"totalTimeElapsed,omitempty"`
	HarvestQuantityRemaining int                    `json:"harvestQuantityRemaining,omitempty"`
	Crop                     collections_crops.Crop `json:"crop,omitempty"`
	CurrentState             int                    `json:"currentState,omitempty"`
	ThiefedBy                []string               `json:"thiefedBy,omitempty"`
	FullyMatured             bool                   `json:"fullyMatured,omitempty"`
	IsPlanted                bool                   `json:"isPlanted,omitempty"`
	IsFertilized             bool                   `json:"isFertilized,omitempty"`
}

type AnimalInfo struct {
	CurrentGrowthTime        int64                      `json:"currentGrowth,omitempty"`
	CurrentHungryTime        int64                      `json:"currentHungryTime,omitempty"`
	CurrentYieldTime         int64                      `json:"currentYieldTime,omitempty"`
	HasYielded               bool                       `json:"hasYielded,omitempty"`
	IsAdult                  bool                       `json:"isAdult,omitempty"`
	Animal                   collections_animals.Animal `json:"animal,omitempty"`
	NeedFed                  bool                       `json:"needFed,omitempty"`
	HarvestQuantityRemaining int                        `json:"harvestQuantityRemaining,omitempty"`
	ThiefedBy                []string                   `json:"thiefedBy,omitempty"`
	AlreadySick              bool                       `json:"alreadySick,omitempty"`
	IsSick                   bool                       `json:"isSick,omitempty"`
}

type BuildingInfo struct {
	CurrentUpgrade int                            `json:"currentUpgrade,omitempty"`
	Occupancy      int                            `json:"currentStageTimeElapsed,omitempty"`
	Building       collections_buildings.Building `json:"building,omitempty"`
}

type PlacedItem struct {
	Key            string         `json:"key,omitempty"`
	ReferenceKey   string         `json:"referenceKey,omitempty"`
	InventoryKey   string         `json:"inventoryKey,omitempty"`
	Position       Position       `json:"position,omitempty"`
	Type           int            `json:"type,omitempty"`
	SeedGrowthInfo SeedGrowthInfo `json:"seedGrowthInfo,omitempty"`
	BuildingInfo   BuildingInfo   `json:"buildingInfo,omitempty"`
	AnimalInfo     AnimalInfo     `json:"animalInfo,omitempty"`
	//for placed items that need parents
	ParentPlacedItemKey string `json:"parentPlacedItemKey,omitempty"`
}
