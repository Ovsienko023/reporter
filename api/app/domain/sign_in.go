package domain

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

type SignInRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (r *SignInRequest) ToDbSignIn() *database.SignIn {
	return &database.SignIn{
		Login: r.Login,
	}
}

type SignInResponse struct {
	Token *string `json:"token,omitempty"`
}
