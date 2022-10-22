package domain

import "github.com/Ovsienko023/reporter/infrastructure/database"

type SignUpRequest struct {
	Login       string  `json:"login,omitempty"`
	Password    string  `json:"password,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

func (r *SignUpRequest) ToDbSignUp() *database.SignUp {
	return &database.SignUp{
		Login:       r.Login,
		Password:    r.Password,
		DisplayName: r.DisplayName,
	}
}
