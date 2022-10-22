package domain

import "github.com/Ovsienko023/reporter/infrastructure/database"

type GetTokenRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (r *GetTokenRequest) ToDbGetToken() *database.GetAuthUser {
	return &database.GetAuthUser{
		Login: r.Login,
	}
}

type GetTokenResponse struct {
	Token *string `json:"token,omitempty"`
}
