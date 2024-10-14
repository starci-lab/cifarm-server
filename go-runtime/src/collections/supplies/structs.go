package collections_supplies

type FertilizerEffect struct {
	TimeReduce int64 `json:"timeReduce"`
}

type Supply struct {
	Key              string           `json:"key"`
	Price            int64            `json:"price"`
	AvailableInShop  bool             `json:"availableInShop"`
	Type             int              `json:"type"`
	FertilizerEffect FertilizerEffect `json:"fertilizerEffect"`
}
