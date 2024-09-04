package collections

type Animal struct {
	Id             string  `json:"id"`
	YieldTime      float32 `json:"yieldTime"`
	OffspringPrice float32 `json:"offspringPrice"`
	Premium        bool    `json:"premium"`
	GrowthTime     float32 `json:"growthTime"`
}
