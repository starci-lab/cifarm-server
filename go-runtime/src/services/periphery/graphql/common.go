package services_periphery_graphql

type NftDataResponse struct {
	TokenId      string              `json:"tokenId"`
	Metadata     NftMetadataResponse `json:"metadata"`
	OwnerAddress string              `json:"ownerAddress"`
}

type NftMetadataResponse struct {
	Image      string `json:"image"`
	Properties string `json:"properties"`
}
