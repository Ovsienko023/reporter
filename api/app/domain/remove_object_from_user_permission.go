package domain

import "github.com/Ovsienko023/reporter/app/repository"

type RemoveObjectFromUserPermissionRequest struct {
	Token    string `json:"token,omitempty" swaggerignore:"true"`
	UserId   string `json:"user_id,omitempty"`
	ObjectId string `json:"object_id,omitempty"`
}

func (r *RemoveObjectFromUserPermissionRequest) ToDbRemoveObjectFromUserPermission(invokerId string) *repository.RemoveObjectFromUserPermission {
	return &repository.RemoveObjectFromUserPermission{
		InvokerId: invokerId,
		UserId:    r.UserId,
		ObjectId:  r.ObjectId,
	}
}
