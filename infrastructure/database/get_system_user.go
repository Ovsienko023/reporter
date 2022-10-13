package database

import (
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

func (c *Client) GetSystemUser() *SystemUser {
	return &SystemUser{
		UserId:      ptr.String("781785ff-676f-46e6-9b6f-e6438d96fe7c"),
		DisplayName: ptr.String("SystemUser"),
	}
}

func (s *ReportLocalStorage) GetSystemUser() *SystemUser {
	return &SystemUser{
		UserId:      ptr.String("11111111-1111-1111-1111-111111111111"),
		DisplayName: ptr.String("SystemUser"),
	}
}

type SystemUser struct {
	UserId      *string `json:"user_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}
