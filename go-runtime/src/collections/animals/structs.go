package collections_animals

type Animal struct {
	Key             string `json:"key"`
	YieldTime       int64  `json:"yieldTime"`
	OffspringPrice  int64  `json:"offspringPrice"`
	Premium         bool   `json:"premium"`
	GrowthTime      int64  `json:"growthTime"`
	AvailableInShop bool   `json:"availableInShop"`
	HungerTime      int64  `json:"hungerTime"`
}
