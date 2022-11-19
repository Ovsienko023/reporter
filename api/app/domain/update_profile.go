package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type UpdateProfileRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DisplayName string `json:"display_name,omitempty"`
}

func (r *UpdateProfileRequest) ToDbUpdateProfile(invokerId string) *repository.UpdateProfile {
	return &repository.UpdateProfile{
		InvokerId:   invokerId,
		DisplayName: r.DisplayName,
	}
}
