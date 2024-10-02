package collections_animals

type Animal struct {
	Key                       string `json:"key"`
	YieldTime                 int64  `json:"yieldTime"`
	OffspringPrice            int64  `json:"offspringPrice"`
	IsNFT                     bool   `json:"isNft"`
	GrowthTime                int64  `json:"growthTime"`
	AvailableInShop           bool   `json:"availableInShop"`
	HungerTime                int64  `json:"hungerTime"`
	MinHarvestQuantity        int    `json:"minHarvestQuantity"`
	MaxHarvestQuantity        int    `json:"maxHarvestQuantity"`
	BasicHarvestExperiences   int64  `json:"basicHarvestExperiences"`
	PremiumHarvestExperiences int64  `json:"premiumHarvestExperiences"`
}
