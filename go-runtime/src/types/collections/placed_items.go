package collections

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type PlacedItem struct {
	Id       string   `json:"id"`
	Position Position `json:"position"`
	Type     int      `json:"type"`
}
