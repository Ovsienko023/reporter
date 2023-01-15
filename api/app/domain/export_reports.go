package domain

import (
	"fmt"
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/safe"
	"strconv"
	"time"
)

type ExportReportsRequest struct {
	Token    string     `json:"token,omitempty" swaggerignore:"true"`
	DateFrom *time.Time `json:"date_from,omitempty"`
	DateTo   *time.Time `json:"date_to,omitempty"`
}

func (r *ExportReportsRequest) ToDbGetReports(invokerId string) *repository.GetReports {
	page := 1
	pageSize := 10000
	return &repository.GetReports{
		InvokerId: invokerId,
		DateFrom:  r.DateFrom,
		DateTo:    r.DateTo,
		Page:      &page,
		PageSize:  &pageSize,
	}
}

type ExportReportsResponse struct {
	Reports []byte
}

func (r *ExportReportsResponse) From(reports []repository.ReportItem) *ExportReportsResponse {
	r.Reports = ToCsvByte(reports)
	return r
}

func ToCsvByte(reports []repository.ReportItem) []byte {
	rawData := "Id,DisplayName,Date,CreatorId,CreatedAt,UpdatedAt,StartTime,EndTime,BreakTime,WorkTime,Body\n"

	for _, report := range reports {
		startTime := "NULL"
		endTime := "NULL"
		breakTime := "NULL"
		workTime := "NULL"

		if report.StartTime != nil {
			startTime = strconv.Itoa(int(*report.StartTime))
		}
		if report.EndTime != nil {
			endTime = strconv.Itoa(int(*report.EndTime))
		}
		if report.BreakTime != nil {
			breakTime = strconv.Itoa(int(*report.BreakTime))
		}
		if report.WorkTime != nil {
			workTime = strconv.Itoa(int(*report.WorkTime))
		}

		rawData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			safe.String(report.Id),
			safe.String(report.DisplayName),
			report.Date.String(),
			safe.String(report.CreatorId),
			report.CreatedAt.String(),
			report.UpdatedAt.String(),
			startTime,
			endTime,
			breakTime,
			workTime,
			safe.String(report.Body),
		)
	}

	return []byte(rawData)
}
