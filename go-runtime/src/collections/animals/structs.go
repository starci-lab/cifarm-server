package collections_animals

type Animal struct {
	Key                       string  `json:"key,omitempty"`
	YieldTime                 int64   `json:"yieldTime,omitempty"`
	OffspringPrice            int64   `json:"offspringPrice,omitempty"`
	IsNFT                     bool    `json:"isNft,omitempty"`
	GrowthTime                int64   `json:"growthTime,omitempty"`
	AvailableInShop           bool    `json:"availableInShop,omitempty"`
	HungerTime                int64   `json:"hungerTime,omitempty"`
	MinHarvestQuantity        int     `json:"minHarvestQuantity,omitempty"`
	MaxHarvestQuantity        int     `json:"maxHarvestQuantity,omitempty"`
	BasicHarvestExperiences   int64   `json:"basicHarvestExperiences,omitempty"`
	PremiumHarvestExperiences int64   `json:"premiumHarvestExperiences,omitempty"`
	Type                      int     `json:"type,omitempty"`
	SickChance                float64 `json:"sickChance,omitempty"`
}
