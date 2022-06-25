package localstore

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/internal/report"
	"github.com/google/uuid"
	"sync"
	"time"
)

type ReportLocalStorage struct {
	reports map[string]*report.Report
	mutex   *sync.Mutex
}

func NewReportLocalStorage() *ReportLocalStorage {
	return &ReportLocalStorage{
		reports: make(map[string]*report.Report),
		mutex:   new(sync.Mutex),
	}
}

func (s *ReportLocalStorage) GetReport(ctx context.Context, msg *report.GetReportRequest) (*report.Report, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if record, ok := s.reports[msg.ReportId]; ok {
		if msg.InvokerId == record.CreatorId {
			return record, nil
		}
	}

	return nil, errors.New("report not found")
}

func (s *ReportLocalStorage) GetReports(ctx context.Context, msg *report.GetReportsRequest) (*report.Reports, error) {
	var reports report.Reports

	s.mutex.Lock()
	for _, val := range s.reports {
		if msg.InvokerId == val.CreatorId {
			reports.Reports = append(reports.Reports, *val)
		}
	}
	s.mutex.Unlock()

	return &reports, nil
}

func (s *ReportLocalStorage) CreateReport(ctx context.Context, msg *report.CreateReportRequest) (*report.CreatedReport, error) {
	s.mutex.Lock()

	id := uuid.New().String()
	s.reports[id] = &report.Report{
		Id:        id,
		Title:     msg.Title,
		CreatorId: msg.InvokerId,
		CreatedAt: int(time.Now().Unix()),
		StartTime: msg.StartTime,
		EndTime:   msg.EndTime,
		BreakTime: msg.BreakTime,
		WorkTime:  msg.WorkTime,
		Template: report.ReportTemplate{
			Variables: msg.Template.Variables,
			Markup:    msg.Template.Markup,
		},
	}
	s.mutex.Unlock()

	record := &report.CreatedReport{
		Id: id,
	}

	return record, nil

}
