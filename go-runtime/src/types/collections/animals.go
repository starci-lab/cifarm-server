package collections

type Animal struct {
	Id             string `json:"id"`
	YieldTime      int64  `json:"yieldTime"`
	OffspringPrice int64  `json:"offspringPrice"`
	Premium        bool   `json:"premium"`
	GrowthTime     int64  `json:"growthTime"`
}
