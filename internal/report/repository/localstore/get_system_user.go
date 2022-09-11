package localstore

import (
	"github.com/Ovsienko023/reporter/internal/report"
	"github.com/Ovsienko023/reporter/pkg/utils/ptr"
)

func (s *ReportLocalStorage) GetSystemUser() *report.SystemUser {
	return &report.SystemUser{
		UserId:      ptr.String("11111111-1111-1111-1111-111111111111"),
		DisplayName: ptr.String("SystemUser"),
	}
}
