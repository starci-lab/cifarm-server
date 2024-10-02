package collections_supplies

type Supply struct {
	Key             string `json:"key"`
	Price           int64  `json:"price"`
	AvailableInShop bool   `json:"availableInShop"`
	Type            int    `json:"type"`
}
