package collections_tiles

type Tile struct {
	ReferenceId  string `json:"referenceId"`
	Price        int64  `json:"price"`
	MaxOwnership int    `json:"maxOwnership"`
}
