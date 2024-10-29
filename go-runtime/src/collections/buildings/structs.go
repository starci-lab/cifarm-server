package collections_buildings

type Building struct {
	Key             string `json:"key,omitempty"`
	AvailableInShop bool   `json:"availableInShop,omitempty"`
	MaxUpgrade      int    `json:"maxUpgrade,omitempty"`
	Price           int64  `json:"price,omitempty"`
	Type            int    `json:"type,omitempty"`
	//key 0 = initial
	Upgrades map[int]Upgrade `json:"upgrades,omitempty"`
}

type Upgrade struct {
	UpgradePrice int64 `json:"upgradePrice,omitempty"`
	Capacity     int   `json:"capacity,omitempty"`
}
