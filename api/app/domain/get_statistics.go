package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"strconv"
	"time"
)

type GetStatisticsRequest struct {
	Token     string  `json:"token,omitempty" swaggerignore:"true"`
	FromDate  int64   `json:"from_date,omitempty"`
	ToDate    int64   `json:"to_date,omitempty"`
	AllowedTo *string `json:"allowed_to,omitempty"`
}

func (r *GetStatisticsRequest) ToDbGetStatistics(invokerId string) *repository.GetStatistics {
	return &repository.GetStatistics{
		InvokerId: invokerId,
		FromDate:  time.Unix(r.FromDate, 0).UTC(),
		ToDate:    time.Unix(r.ToDate, 0).UTC(),
		AllowedTo: r.AllowedTo,
	}
}

func ToGetStatisticsRequest(token, fromDate, toData, allowedTo string) *GetStatisticsRequest {
	// todo Подумать над добавлением уровня сообщений для transport.http и добавлением валидации

	from, _ := strconv.ParseInt(fromDate, 10, 64)
	to, _ := strconv.ParseInt(toData, 10, 64)

	res := &GetStatisticsRequest{
		Token:    token,
		FromDate: from,
		ToDate:   to,
	}

	if allowedTo != "" {
		res.AllowedTo = &allowedTo
	}

	return res
}

type GetStatisticsResponse struct {
	AvgHoursWorked *int64 `json:"avg_hours_worked,omitempty"`
	HoursWorked    *int64 `json:"hours_worked,omitempty"`
	AvgHoursBreak  *int64 `json:"avg_hours_break,omitempty"`
	AvgStartTime   *int64 `json:"avg_start_time,omitempty"`
}

func FromGetStatisticsResponse(stats *repository.Statistics) *GetStatisticsResponse {
	if stats == nil {
		return nil
	}

	return &GetStatisticsResponse{
		AvgHoursWorked: stats.AvgHoursWorked,
		HoursWorked:    stats.HoursWorked,
		AvgHoursBreak:  stats.AvgHoursBreak,
		AvgStartTime:   stats.AvgStartTime,
	}
}
