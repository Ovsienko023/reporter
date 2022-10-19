package domain

type GetTokenRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetTokenResponse struct {
	Token *string `json:"token,omitempty"`
}
