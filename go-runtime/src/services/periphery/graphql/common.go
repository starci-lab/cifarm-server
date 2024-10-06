package services_periphery_graphql

type NftDataResponse struct {
	TokenId      string `json:"tokenId"`
	TokenURI     string `json:"tokenURI"`
	OwnerAddress string `json:"ownerAddress"`
}
