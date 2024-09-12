package collections_inventories

type Inventory struct {
	Key          string `json:"key"`
	ReferenceKey string `json:"referenceKey"`
	Type         int    `json:"type"`
	Quantity     int    `json:"quantity"`
}
