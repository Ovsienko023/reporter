package database

import (
	"context"
)

func (s *ReportLocalStorage) DeleteReport(ctx context.Context, msg *DeleteReport) error {
	s.mutex.Lock()

	if _, ok := s.reports[msg.ReportId]; ok {
		delete(s.reports, msg.ReportId)
	} else {
		return ErrReportIdNotFound
	}

	s.mutex.Unlock()
	return nil
}

type DeleteReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"id,omitempty"`
}
