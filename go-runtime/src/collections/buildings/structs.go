package collections_buildings

type Building struct {
	Key             string `json:"key"`
	AvailableInShop bool   `json:"availableInShop"`
	MaxUpgrade      int    `json:"maxUpgrade"`
	Price           int64  `json:"price"`
	Capacity        int    `json:"capacity"`
	AnimalKey       string `json:"animalKey"`
}
