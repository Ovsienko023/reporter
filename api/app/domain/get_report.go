package domain

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetReportRequest struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

func (r *GetReportRequest) ToDbGetReport() *database.GetReport {
	return &database.GetReport{
		InvokerId: r.InvokerId,
		ReportId:  r.ReportId,
	}
}

type GetReportResponse struct {
	Report *Report `json:"report,omitempty"`
}

type Report struct {
	Id        *string `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	Date      *int64  `json:"date,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	StartTime *int64  `json:"start_time,omitempty"`
	EndTime   *int64  `json:"end_time,omitempty"`
	BreakTime *int64  `json:"break_time,omitempty"`
	WorkTime  *int64  `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

func FromGetReportResponse(resp *database.Report) *GetReportResponse {
	if resp == nil {
		return nil
	}

	return &GetReportResponse{
		Report: &Report{
			Id:        resp.Id,
			Title:     resp.Title,
			Date:      ptr.Int64(resp.Date.Unix()),
			CreatorId: resp.CreatorId,
			CreatedAt: ptr.Int64(resp.CreatedAt.Unix()),
			UpdatedAt: ptr.Int64(resp.UpdatedAt.Unix()),
			StartTime: ptr.Int64(resp.StartTime.Unix()),
			EndTime:   ptr.Int64(resp.EndTime.Unix()),
			BreakTime: ptr.Int64(resp.BreakTime.Unix()),
			WorkTime:  ptr.Int64(resp.WorkTime.Unix()),
			Body:      resp.Body,
		},
	}
}
