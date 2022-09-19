package database

import (
	"context"
	"github.com/google/uuid"
	"time"
)

func (c *Client) CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error) {
	return nil, nil
}

func (s *ReportLocalStorage) CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error) {
	s.mutex.Lock()

	id := uuid.New().String()
	createdAt := int(time.Now().Unix())
	s.reports[id] = &Report{
		Id:        &id,
		Title:     &msg.Title,
		CreatorId: &msg.InvokerId,
		CreatedAt: &createdAt,
		StartTime: &msg.StartTime,
		EndTime:   &msg.EndTime,
		BreakTime: &msg.BreakTime,
		WorkTime:  &msg.WorkTime,
		Body:      &msg.Body,
	}
	s.mutex.Unlock()

	record := &CreatedReport{
		Id: id,
	}

	return record, nil
}

type CreateReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	Title     string `json:"title,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type CreatedReport struct {
	Id string `json:"id,omitempty"`
}
