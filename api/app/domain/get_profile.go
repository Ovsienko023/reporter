package domain

import "github.com/Ovsienko023/reporter/infrastructure/database"

type GetProfileRequest struct {
	Token string `json:"token,omitempty" swaggerignore:"true"`
}

func (r *GetProfileRequest) ToDbGetProfile(invokerId string) *database.GetProfile {
	return &database.GetProfile{
		InvokerId: invokerId,
	}
}

type GetProfileResponse struct {
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Login       *string `json:"login,omitempty"`
}

func FromGetProfileResponse(resp *database.Profile) *GetProfileResponse {
	if resp == nil {
		return nil
	}

	return &GetProfileResponse{
		Id:          resp.Id,
		DisplayName: resp.DisplayName,
		Login:       resp.Login,
	}
}
