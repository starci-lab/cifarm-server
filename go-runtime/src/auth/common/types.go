package common

type VerifyMessageRequestBody struct {
	Message   string  `json:"message"`
	Signature string  `json:"signature"`
	PublicKey string  `json:"public_key"`
	Platform  *string `json:"platform,omitempty"`
}

type VerifyMessageResponseData struct {
	Result  bool   `json:"result"`
	Address string `json:"address"`
}
