package domain

import "github.com/Ovsienko023/reporter/infrastructure/database"

type DeleteReportRequest struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

func (r *DeleteReportRequest) ToDbDeleteReport() *database.DeleteReport {
	return &database.DeleteReport{
		InvokerId: r.InvokerId,
		ReportId:  r.ReportId,
	}
}
