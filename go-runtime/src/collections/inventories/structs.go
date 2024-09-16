package collections_inventories

type TileInfo struct {
	GrowthTimeReduction  int `json:"growthTimeReduction"`
	PestResistance       int `json:"pestResistance"`
	ProductivityIncrease int `json:"productivityIncrease"`
	WeedResistance       int `json:"weedResistance"`
}

type Inventory struct {
	Key          string   `json:"key"`
	ReferenceKey string   `json:"referenceKey"`
	TileInfo     TileInfo `json:"tileInfo"`
	Type         int      `json:"type"`
	Quantity     int      `json:"quantity"`
	Unique       bool     `json:"unique"`
	TokenId      int      `json:"tokenId"`
	Placeable    bool     `json:"placeable"`
	IsPlaced     bool     `json:"isPlaced"`
	IsPremium    bool     `json:"isPremium"`
	Deliverable  bool     `json:"deliverable"`
}
