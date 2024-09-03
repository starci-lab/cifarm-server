package collections

type Inventory struct {
	Id       string `json:"id"`
	Type     int    `json:"type"`
	Quantity int    `json:"quantity"`
}

const (
	TYPE_SEED = 0
)

const (
	INVENTORY_CARROT_SEED = 0
	INVENTORY_POTATO_SEED = 0
)
