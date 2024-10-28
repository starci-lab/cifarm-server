package services_periphery_graphql

type NftDataResponse struct {
	TokenId      string              `json:"tokenId,omitempty"`
	Metadata     NftMetadataResponse `json:"metadata,omitempty"`
	OwnerAddress string              `json:"ownerAddress,omitempty"`
}

type NftMetadataResponse struct {
	Image      string `json:"image,omitempty"`
	Properties string `json:"properties,omitempty"`
}
