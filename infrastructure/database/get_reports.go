package database

import (
	"context"
)

func (s *ReportLocalStorage) GetReports(ctx context.Context, msg *GetReports) (*Reports, *int, error) {
	var reports Reports

	s.mutex.Lock()
	for _, val := range s.reports {
		if msg.InvokerId == *val.CreatorId {
			reports.Reports = append(reports.Reports, *val)
		}
	}
	s.mutex.Unlock()

	count := len(reports.Reports)

	return &reports, &count, nil
}

type GetReports struct {
	InvokerId string `json:"invoker_id,omitempty"`
}

type Reports struct {
	Reports []Report `json:"reports,omitempty"`
}
