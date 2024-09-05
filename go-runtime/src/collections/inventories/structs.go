package collections_inventories

type Inventory struct {
	ReferenceId string `json:"referenceId"`
	Type        int    `json:"type"`
	Quantity    int    `json:"quantity"`
}
