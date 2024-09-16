package collections_config

type Metadata struct {
	Key            string `json:"key"`
	ChainKey       string `json:"chainKey"`
	AccountAddress string `json:"accountAddress"`
	Network        string `json:"network"`
}

type VisitState struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}
