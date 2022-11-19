package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type SignInRequest struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (r *SignInRequest) ToDbSignIn() *repository.SignIn {
	return &repository.SignIn{
		Login: r.Login,
	}
}

type SignInResponse struct {
	Token *string `json:"token,omitempty"`
}
