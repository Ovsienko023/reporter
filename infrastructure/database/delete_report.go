package database

import (
	"context"
	"errors"
)

const sqlDeleteReport = `
	update main.reports
    set deleted_at = now()
    where id = $1
    returning 1;`

func (c *Client) DeleteReport(ctx context.Context, msg *DeleteReport) error {
	err := c.isFindReport(ctx, msg.ReportId)
	if errors.Is(err, ErrReportIdNotFound) {
		return ErrReportIdNotFound
	}
	//_, err := c.GetReport(ctx, &GetReport{ReportId: msg.ReportId})
	//if errors.Is(err, ErrReportIdNotFound) {
	//	return ErrReportIdNotFound
	//}

	row, err := c.driver.Query(ctx, sqlDeleteReport, msg.ReportId)
	if err != nil {
		return NewInternalError(err)
	}

	var deleted *int

	row.Next()
	err = row.Scan(&deleted)
	if err != nil {
		return NewInternalError(err)
	}

	return nil
}

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
