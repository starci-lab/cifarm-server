package collections_spin

type Spin struct {
	Key string `json:"key"`

	// quantity of reward
	GoldAmount  int64   `json:"goldAmount"`
	TokenAmount float64 `json:"tokenAmount"`
	Quantity    int     `json:"quantity"`

	Type int `json:"type"`
	//to indecate if random 100, the range threshold acheive the reward
	ThresholdMin float64 `json:"thresholdMin"`
	ThresholdMax float64 `json:"thresholdMax"`
}
