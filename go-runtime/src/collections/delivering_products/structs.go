package collections_delivering_products

type DeliveringProduct struct {
	Key          string `json:"key"`
	ReferenceKey string `json:"referenceKey"`
	Type         int    `json:"type"`
	Quantity     int    `json:"quantity"`
	IsPremium    bool   `json:"isPremium"`
}
