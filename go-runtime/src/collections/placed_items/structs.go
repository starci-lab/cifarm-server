package collections_placed_items

import (
	collections_animals "cifarm-server/src/collections/animals"
	collections_buildings "cifarm-server/src/collections/buildings"
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
	FullyMatured             bool                   `json:"fullyMatured"`
	IsPlanted                bool                   `json:"isPlanted"`
}

type AnimalInfo struct {
	CurrentGrowthTime        int64                      `json:"currentGrowth"`
	CurrentHungryTime        int64                      `json:"currentHungryTime"`
	CurrentYieldTime         int64                      `json:"currentYieldTime"`
	HasYielded               bool                       `json:"hasYielded"`
	IsAdult                  bool                       `json:"isAdult"`
	Animal                   collections_animals.Animal `json:"animal"`
	NeedFed                  bool                       `json:"needFed"`
	HarvestQuantityRemaining int                        `json:"harvestQuantityRemaining"`
	ThiefedBy                []string                   `json:"thiefedBy"`
}

type BuildingInfo struct {
	CurrentUpgrade int                            `json:"currentUpgrade"`
	Occupancy      int                            `json:"currentStageTimeElapsed"`
	Building       collections_buildings.Building `json:"building"`
}

type PlacedItem struct {
	Key            string         `json:"key"`
	ReferenceKey   string         `json:"referenceKey"`
	InventoryKey   string         `json:"inventoryKey"`
	Position       Position       `json:"position"`
	Type           int            `json:"type"`
	SeedGrowthInfo SeedGrowthInfo `json:"seedGrowthInfo"`
	BuildingInfo   BuildingInfo   `json:"buildingInfo"`
	AnimalInfo     AnimalInfo     `json:"animalInfo"`
	//for placed items that need parents
	ParentPlacedItemKey string `json:"parentPlacedItemKey"`
}
