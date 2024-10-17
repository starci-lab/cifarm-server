package collections_buildings

type Building struct {
	Key             string `json:"key"`
	AvailableInShop bool   `json:"availableInShop"`
	MaxUpgrade      int    `json:"maxUpgrade"`
	Price           int64  `json:"price"`
	Type            int    `json:"type"`
	//key 0 = initial
	Upgrades map[int]Upgrade `json:"upgrades"`
}

type Upgrade struct {
	UpgradePrice int64 `json:"upgradePrice"`
	Capacity     int   `json:"capacity"`
}
