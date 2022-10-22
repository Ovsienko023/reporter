package domain

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetReportsRequest struct {
	Token string `json:"token,omitempty" swaggerignore:"true"`
}

func (r *GetReportsRequest) ToDbGetReports(invokerId string) *database.GetReports {
	return &database.GetReports{
		InvokerId: invokerId,
	}
}

type GetReportsResponse struct {
	Count   *int         `json:"count,omitempty"`
	Reports []ReportItem `json:"reports" json:"reports,omitempty"`
}

type ReportItem struct {
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

func FromGetReportsResponse(resp []database.ReportItem, cnt *int) *GetReportsResponse {
	if resp == nil {
		return nil
	}

	reports := make([]ReportItem, 0, len(resp))

	for _, obj := range resp {
		item := ReportItem{
			Id:        obj.Id,
			Title:     obj.Title,
			Date:      ptr.Int64(obj.Date.Unix()),
			CreatorId: obj.CreatorId,
			CreatedAt: ptr.Int64(obj.CreatedAt.Unix()),
			UpdatedAt: ptr.Int64(obj.UpdatedAt.Unix()),
			StartTime: ptr.Int64(obj.StartTime.Unix()),
			EndTime:   ptr.Int64(obj.EndTime.Unix()),
			BreakTime: ptr.Int64(obj.BreakTime.Unix()),
			WorkTime:  ptr.Int64(obj.WorkTime.Unix()),
			Body:      obj.Body,
		}
		reports = append(reports, item)
	}

	return &GetReportsResponse{
		Count:   cnt,
		Reports: reports,
	}
}
