package collections_animals

type Animal struct {
	ReferenceId    string `json:"referenceId"`
	YieldTime      int64  `json:"yieldTime"`
	OffspringPrice int64  `json:"offspringPrice"`
	Premium        bool   `json:"premium"`
	GrowthTime     int64  `json:"growthTime"`
}
