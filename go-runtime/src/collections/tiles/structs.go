package collections_tiles

type Tile struct {
	Key             string `json:"key,omitempty"`
	Price           int64  `json:"price,omitempty"`
	MaxOwnership    int    `json:"maxOwnership,omitempty"`
	IsNFT           bool   `json:"isNft,omitempty"`
	AvailableInShop bool   `json:"availableInShop,omitempty"`
}
