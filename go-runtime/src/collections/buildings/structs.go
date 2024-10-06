package collections_buildings

type Building struct {
	Key              string                 `json:"key"`
	AvailableInShop  bool                   `json:"availableInShop"`
	MaxUpgrade       int                    `json:"maxUpgrade"`
	UpgradeSummaries map[int]UpgradeSummary `json:"upgradeSummaries"`
	AnimalKey        string                 `json:"animalKey"`
}

type UpgradeSummary struct {
	Price    int64 `json:"price"`
	Capacity int   `json:"capacity"`
}
