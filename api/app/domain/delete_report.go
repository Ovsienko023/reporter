package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteReportRequest struct {
	Token    string `json:"token,omitempty"`
	ReportId string `json:"report_id,omitempty"`
}

func (r *DeleteReportRequest) ToDbDeleteReport(invokerId string) *repository.DeleteReport {
	return &repository.DeleteReport{
		InvokerId: invokerId,
		ReportId:  r.ReportId,
	}
}
