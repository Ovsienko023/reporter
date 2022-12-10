package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetReportRequest struct {
	Token    string `json:"token,omitempty" swaggerignore:"true"`
	ReportId string `json:"report_id,omitempty" swaggerignore:"true"`
}

func (r *GetReportRequest) ToDbGetReport(invokerId string) *repository.GetReport {
	return &repository.GetReport{
		InvokerId: invokerId,
		ReportId:  r.ReportId,
	}
}

type GetReportResponse struct {
	Report *Report `json:"report,omitempty"`
}

type Report struct {
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Date        *int64  `json:"date,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	CreatedAt   *int64  `json:"created_at,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	StartTime   *int64  `json:"start_time,omitempty"`
	EndTime     *int64  `json:"end_time,omitempty"`
	BreakTime   *int64  `json:"break_time,omitempty"`
	WorkTime    *int64  `json:"work_time,omitempty"`
	Body        *string `json:"body,omitempty"`
}

func FromGetReportResponse(resp *repository.Report) *GetReportResponse {
	if resp == nil {
		return nil
	}

	return &GetReportResponse{
		Report: &Report{
			Id:          resp.Id,
			DisplayName: resp.DisplayName,
			Date:        ptr.Int64(resp.Date.Unix()),
			CreatorId:   resp.CreatorId,
			CreatedAt:   ptr.Int64(resp.CreatedAt.Unix()),
			UpdatedAt:   ptr.Int64(resp.UpdatedAt.Unix()),
			StartTime:   resp.StartTime,
			EndTime:     resp.EndTime,
			BreakTime:   resp.BreakTime,
			WorkTime:    resp.WorkTime,
			Body:        resp.Body,
		},
	}
}
