package collections_crops

type Crop struct {
	Key                         string `json:"key"`
	GrowthStageDuration         int64  `json:"growthStageDuration"`
	GrowthStages                int    `json:"growthStages"`
	Price                       int64  `json:"price"`
	Premium                     bool   `json:"premium"`
	Perennial                   bool   `json:"perennial"`
	NextGrowthStageAfterHarvest int    `json:"nextGrowthStageAfterHarvest"`
	MinHarvestQuantity          int    `json:"minHarvestQuantity"`
	MaxHarvestQuantity          int    `json:"maxHarvestQuantity"`
	BasicHarvestExperiences     int    `json:"basicHarvestExperiences"`
	PremiumHarvestExperiences   int    `json:"premiumHarvestExperiences"`
	AvailableInShop             bool   `json:"availableInShop"`
}
