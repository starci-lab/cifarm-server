package collections_delivering_products

type DeliveringProduct struct {
	Key          string `json:"key,omitempty"`
	ReferenceKey string `json:"referenceKey,omitempty"`
	Type         int    `json:"type,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	Premium      bool   `json:"premium,omitempty"`
	Index        int    `json:"index,omitempty"`
}
