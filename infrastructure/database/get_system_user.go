package database

import (
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

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
