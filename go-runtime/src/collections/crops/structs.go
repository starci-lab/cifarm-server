package collections_crops

type Crop struct {
	Key                         string `json:"key,omitempty"`
	GrowthStageDuration         int64  `json:"growthStageDuration,omitempty"`
	GrowthStages                int    `json:"growthStages,omitempty"`
	Price                       int64  `json:"price,omitempty"`
	Premium                     bool   `json:"premium,omitempty"`
	Perennial                   bool   `json:"perennial,omitempty"`
	NextGrowthStageAfterHarvest int    `json:"nextGrowthStageAfterHarvest,omitempty"`
	MinHarvestQuantity          int    `json:"minHarvestQuantity,omitempty"`
	MaxHarvestQuantity          int    `json:"maxHarvestQuantity,omitempty"`
	BasicHarvestExperiences     int    `json:"basicHarvestExperiences,omitempty"`
	PremiumHarvestExperiences   int    `json:"premiumHarvestExperiences,omitempty"`
	AvailableInShop             bool   `json:"availableInShop,omitempty"`
}
