package collections

type PlantSeed struct {
	Id                          string  `json:"id"`
	GrowthStageDuration         float32 `json:"growthStageDuration"`
	GrowthStages                int     `json:"growthStages"`
	SeedPrice                   float32 `json:"price"`
	Premium                     bool    `json:"premium"`
	Perennial                   bool    `json:"perennial"`
	NextGrowthStageAfterHarvest int     `json:"nextGrowthStageAfterHarvest"`
	MinHarvestQuantity          int     `json:"minHarvestQuantity"`
	MaxHarvestQuantity          int     `json:"maxHarvestQuantity"`
}

const (
	PLANT_SEED_CARROT_SEED = "carrotSeed"
	PLANT_SEED_POTATO_SEED = "potatoSeed"
)
