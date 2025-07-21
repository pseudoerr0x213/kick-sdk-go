package kick

type GetPublicKeyResponse struct {
	Data    PublicKey `json:"data"`
	Message string    `json:"message"`
}

type PublicKey struct {
	PublicKey string `json:"public_key"`
}
