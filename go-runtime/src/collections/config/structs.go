package collections_config

type Metadata struct {
	ChainKey       string `json:"chainKey"`
	AccountAddress string `json:"accountAddress"`
	Network        string `json:"network"`
}

type VisitState struct {
	UserId string `json:"userId"`
}
