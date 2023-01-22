package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type GetUsersRequest struct {
	Token    string `json:"token,omitempty" swaggerignore:"true"`
	Page     *int   `json:"page,omitempty"`
	PageSize *int   `json:"page_size,omitempty"`
}

func (r *GetUsersRequest) ToDbGetUsers(invokerId string) *repository.GetUsers {
	return &repository.GetUsers{
		InvokerId: invokerId,
		Page:      r.Page,
		PageSize:  r.PageSize,
	}
}

type GetUsersResponse struct {
	Count *int        `json:"count,omitempty"`
	Users []UsersItem `json:"reports" json:"reports,omitempty"`
}

type UsersItem struct {
	Id          *string `json:"id,omitempty"`
	Login       *string `json:"login,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

func FromGetUsersResponse(resp []repository.UserItem, cnt *int) *GetUsersResponse {
	if resp == nil {
		return nil
	}

	users := make([]UsersItem, 0, len(resp))

	for _, obj := range resp {
		item := UsersItem{
			Id:          obj.Id,
			DisplayName: obj.DisplayName,
			Login:       obj.Login,
		}
		users = append(users, item)
	}

	return &GetUsersResponse{
		Count: cnt,
		Users: users,
	}
}
