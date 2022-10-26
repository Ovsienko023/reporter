package domain

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

type UpdateProfileRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DisplayName string `json:"display_name,omitempty"`
}

func (r *UpdateProfileRequest) ToDbUpdateProfile(invokerId string) *database.UpdateProfile {
	return &database.UpdateProfile{
		InvokerId:   invokerId,
		DisplayName: r.DisplayName,
	}
}
