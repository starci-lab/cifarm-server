package collections

type PlayerMetadata struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

type Users struct {
	UserIds []string `json:"userIds"`
}
