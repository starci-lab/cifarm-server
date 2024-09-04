package collections

type PlayerMetadata struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type Users struct {
	Id      string   `json:"id"`
	UserIds []string `json:"userIds"`
}
