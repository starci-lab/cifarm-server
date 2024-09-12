package services_periphery_graphql

type NftData struct {
	TokenId      int    `json:"tokenId"`
	TokenURI     string `json:"tokenURI"`
	OwnerAddress string `json:"ownerAddress"`
}
