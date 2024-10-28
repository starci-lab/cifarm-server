package collections_spin

type Spin struct {
	Key string `json:"key,omitempty"`

	// quantity of reward
	GoldAmount  int64   `json:"goldAmount,omitempty"`
	TokenAmount float64 `json:"tokenAmount,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`

	Type int `json:"type,omitempty"`
	//to indecate if random 100, the range threshold acheive the reward
	ThresholdMin float64 `json:"thresholdMin,omitempty"`
	ThresholdMax float64 `json:"thresholdMax,omitempty"`
}
