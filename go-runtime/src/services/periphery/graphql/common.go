package services_periphery_graphql

type NftDataResponse struct {
	TokenId      string `json:"tokenId"`
	TokenURI     string `json:"tokenURI"`
	OwnerAddress string `json:"ownerAddress"`
}

type NftMetadataResponse struct {
	Image      string `json:"image"`
	Properties string `json:"properties"`
}
