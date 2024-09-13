package services_periphery_graphql

type NftDataResponse struct {
	TokenId      int    `json:"tokenId"`
	TokenURI     string `json:"tokenURI"`
	OwnerAddress string `json:"ownerAddress"`
}
