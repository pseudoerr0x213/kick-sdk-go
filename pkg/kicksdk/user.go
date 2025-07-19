package kicksdk

type UserReponse struct {
	Active    bool   `json:"active"`
	ClientID  string `json:"client_id"`
	Exp       int64  `json:"exp"`
	Scope     string `json:"scope"`
	TokenType string `json:"token_type"`
}
