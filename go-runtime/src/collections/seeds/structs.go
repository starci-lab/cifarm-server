package collections_seeds

type Seed struct {
	Key                         string `json:"key"`
	GrowthStageDuration         int64  `json:"growthStageDuration"`
	GrowthStages                int    `json:"growthStages"`
	Price                       int64  `json:"price"`
	Premium                     bool   `json:"premium"`
	Perennial                   bool   `json:"perennial"`
	NextGrowthStageAfterHarvest int    `json:"nextGrowthStageAfterHarvest"`
	MinHarvestQuantity          int    `json:"minHarvestQuantity"`
	MaxHarvestQuantity          int    `json:"maxHarvestQuantity"`
}
