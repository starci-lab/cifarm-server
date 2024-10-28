package collections_supplies

type FertilizerEffect struct {
	TimeReduce int64 `json:"timeReduce,omitempty"`
}

type Supply struct {
	Key              string           `json:"key,omitempty"`
	Price            int64            `json:"price,omitempty"`
	AvailableInShop  bool             `json:"availableInShop,omitempty"`
	Type             int              `json:"type,omitempty"`
	FertilizerEffect FertilizerEffect `json:"fertilizerEffect,omitempty"`
}
