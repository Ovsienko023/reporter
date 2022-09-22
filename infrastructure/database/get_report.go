package database

import (
	"context"
)

func (s *ReportLocalStorage) GetReport(ctx context.Context, msg *GetReport) (*Report, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if record, ok := s.reports[msg.ReportId]; ok {
		if msg.InvokerId == *record.CreatorId {
			return record, nil
		}
	}

	return nil, ErrReportIdNotFound
}

type GetReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

type Report struct {
	Id        *string `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	Date      *int    `json:"date,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	CreatedAt *int    `json:"created_at,omitempty"`
	UpdatedAt *int    `json:"updated_at,omitempty"`
	DeletedAt *int    `json:"deleted_at,omitempty"`
	StartTime *int    `json:"start_time,omitempty"`
	EndTime   *int    `json:"end_time,omitempty"`
	BreakTime *int    `json:"break_time,omitempty"`
	WorkTime  *int    `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}
