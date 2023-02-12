package domain

import "github.com/Ovsienko023/reporter/app/repository"

type AddObjectToUserPermissionRequest struct {
	Token      string `json:"token,omitempty" swaggerignore:"true"`
	UserId     string `json:"user_id,omitempty" swaggerignore:"true"`
	ObjectType string `json:"object_type,omitempty"`
	ObjectId   string `json:"object_id,omitempty"`
}

func (r *AddObjectToUserPermissionRequest) ToDbAddObjectToUserPermission(invokerId string) *repository.AddObjectToUserPermission {
	return &repository.AddObjectToUserPermission{
		InvokerId:  invokerId,
		UserId:     r.UserId,
		ObjectType: r.ObjectType,
		ObjectId:   r.ObjectId,
	}
}
