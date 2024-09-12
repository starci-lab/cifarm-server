package collections_nfts

type Nft struct {
	Key            string `json:"key"`
	Type           int    `json:"type"`
	TokenId        int    `json:"tokenId"`
	AccountAddress string `json:"accountAddress"`
	ChainKey       string `json:"chainKey"`
	Network        string `json:"network"`
}
