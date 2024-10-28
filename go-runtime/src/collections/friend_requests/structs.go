package collections_friend_requests

type TileInfo struct {
	GrowthTimeReduction  int `json:"growthTimeReduction,omitempty"`
	PestResistance       int `json:"pestResistance,omitempty"`
	ProductivityIncrease int `json:"productivityIncrease,omitempty"`
	WeedResistance       int `json:"weedResistance,omitempty"`
}

type Inventory struct {
	Key          string   `json:"key,omitempty"`
	ReferenceKey string   `json:"referenceKey,omitempty"`
	TileInfo     TileInfo `json:"tileInfo,omitempty"`
	Type         int      `json:"type,omitempty"`
	Quantity     int      `json:"quantity,omitempty"`
	Unique       bool     `json:"unique,omitempty"`
	TokenId      string   `json:"tokenId,omitempty"`
	Placeable    bool     `json:"placeable,omitempty"`
	IsPlaced     bool     `json:"isPlaced,omitempty"`
	Premium      bool     `json:"premium,omitempty"`
	Deliverable  bool     `json:"deliverable,omitempty"`
	AsTool       bool     `json:"asTool,omitempty"`
}
