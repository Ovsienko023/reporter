package database

import (
	"context"
	"time"
)

func (s *ReportLocalStorage) UpdateReport(ctx context.Context, msg *UpdateReport) error {
	s.mutex.Lock()

	if _, ok := s.reports[msg.ReportId]; ok {
		createdAt := int(time.Now().Unix())
		s.reports[msg.ReportId] = &Report{
			Id:        &msg.ReportId,
			Title:     msg.Title,
			Date:      msg.Date,
			CreatorId: &msg.InvokerId,
			CreatedAt: &createdAt,
			StartTime: msg.StartTime,
			EndTime:   msg.EndTime,
			BreakTime: msg.BreakTime,
			WorkTime:  msg.WorkTime,
			Body:      msg.Body,
		}
	} else {
		return ErrReportIdNotFound
	}

	s.mutex.Unlock()
	return nil
}

type UpdateReport struct {
	InvokerId string  `json:"invoker_id,omitempty"`
	ReportId  string  `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	Date      *int    `json:"date,omitempty"`
	StartTime *int    `json:"start_time,omitempty"`
	EndTime   *int    `json:"end_time,omitempty"`
	BreakTime *int    `json:"break_time,omitempty"`
	WorkTime  *int    `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}
