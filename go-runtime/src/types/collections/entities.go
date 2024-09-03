package collections

type Plant struct {
	Id                          int     `json:"id"`
	Key                         string  `json:"-"`
	GrowthStageDuration         float32 `json:"growthStageDuration"`
	GrowthStages                int     `json:"growthStages"`
	SeedPrice                   float32 `json:"price"`
	Premium                     bool    `json:"premium"`
	Perennial                   bool    `json:"perennial"`
	NextGrowthStageAfterHarvest int     `json:"nextGrowthStageAfterHarvest"`
	MinHarvestQuantity          int     `json:"minHarvestQuantity"`
	MaxHarvestQuantity          int     `json:"maxHarvestQuantity"`
}

type Animal struct {
	Id             int     `json:"id"`
	Key            string  `json:"-"`
	YieldTime      float32 `json:"yieldTime"`
	OffspringPrice float32 `json:"offspringPrice"`
	Premium        bool    `json:"premium"`
	GrowthTime     float32 `json:"growthTime"`
}
