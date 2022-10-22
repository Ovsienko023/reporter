package domain

import "github.com/Ovsienko023/reporter/infrastructure/database"

type DeleteReportRequest struct {
	Token    string `json:"token,omitempty"`
	ReportId string `json:"report_id,omitempty"`
}

func (r *DeleteReportRequest) ToDbDeleteReport(invokerId string) *database.DeleteReport {
	return &database.DeleteReport{
		InvokerId: invokerId,
		ReportId:  r.ReportId,
	}
}
