package collections_tiles

type Tile struct {
	Key             string `json:"key"`
	Price           int64  `json:"price"`
	MaxOwnership    int    `json:"maxOwnership"`
	IsNFT           bool   `json:"isNft"`
	AvailableInShop bool   `json:"availableInShop"`
}
