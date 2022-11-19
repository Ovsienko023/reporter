package repository

import (
	"context"
	"time"
)

const (
	sqlMostEarlyStartTime = `
	select date, start_time 
	from main.reports 
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
    where rtu.user_id = $1 and 
	      date_trunc('day', date) >= date_trunc('day', $2::timestamp) and 
	      date_trunc('day', date) <= date_trunc('day', $3::timestamp)
	order by start_time limit 1`  // самое раннее начало

	sqlMostLateStartTime = `
	select date, start_time 
	from main.reports 
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
    where rtu.user_id = $1 and 
	      date_trunc('day', date) >= date_trunc('day', $2::timestamp) and 
	      date_trunc('day', date) <= date_trunc('day', $3::timestamp)
	order by start_time desc limit 1`  // самое подзнее начало

	sqlMostShortBreakTime = `
	select date, break_time 
	from main.reports
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
    where rtu.user_id = $1 and 
	      date_trunc('day', date) >= date_trunc('day', $2::timestamp) and 
	      date_trunc('day', date) <= date_trunc('day', $3::timestamp)
	order by break_time limit 1`  // самый короткий перерыв

	sqlMostLongBreakTime = `
	select date, break_time 
	from main.reports
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
    where rtu.user_id = $1 and 
	      date_trunc('day', date) >= date_trunc('day', $2::timestamp) and 
	      date_trunc('day', date) <= date_trunc('day', $3::timestamp)
	order by break_time desc limit 1`  // самыей длинный перерыв

	sqlMostLongDay = `
	select date, (extract(epoch from end_time) -
              extract(epoch from start_time) -
              extract(epoch from break_time)) / 3600 -- самый длинный рабочий день(последний)
	from main.reports
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
	where rtu.user_id = $1 and 
	      (extract(epoch from end_time) -
       	   extract(epoch from start_time) -
           extract(epoch from break_time)) = (select max(extract(epoch from end_time) -
                                                         extract(epoch from start_time) -
                                                         extract(epoch from break_time))
                                             from main.reports)
	order by date desc limit 1`

	sqlGetStatistics = `
	select avg(extract(epoch from end_time) -
               extract(epoch from start_time) -
               extract(epoch from break_time)) as avg_hours_worked,  -- среднне количество отработанных часов
       	   sum(extract(epoch from end_time) -
               extract(epoch from start_time) -
               extract(epoch from break_time)) as hours_worked,      -- отработаннно часов
       	   avg(extract(epoch from break_time)) as avg_hours_break,   -- среднне количество часов перерыва
           avg(extract(epoch from start_time)) as avg_start_time     -- среднее время начала дня
	from main.reports
	inner join main.reports_to_users rtu on reports.id = rtu.report_id
	where rtu.user_id = $1 and 
	      date_trunc('day', date) >= date_trunc('day', $2::timestamp) and 
	      date_trunc('day', date) <= date_trunc('day', $3::timestamp)`
)

func (c *Client) GetStatistics(ctx context.Context, msg *GetStatistics) (*Statistics, error) {
	row, err := c.driver.Query(ctx, sqlGetStatistics,
		msg.InvokerId,
		msg.FromDate,
		msg.ToDate,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	stats := &Statistics{}

	for row.Next() {
		err := row.Scan(
			&stats.AvgHoursWorked,
			&stats.HoursWorked,
			&stats.AvgHoursBreak,
			&stats.AvgStartTime,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	return stats, nil
}

type GetStatistics struct {
	InvokerId string    `json:"invoker_id,omitempty"`
	FromDate  time.Time `json:"from_date,omitempty"`
	ToDate    time.Time `json:"to_date,omitempty"`
}

type Statistics struct {
	AvgHoursWorked *int64 `json:"avg_hours_worked,omitempty"`
	HoursWorked    *int64 `json:"hours_worked,omitempty"`
	AvgHoursBreak  *int64 `json:"avg_hours_break,omitempty"`
	AvgStartTime   *int64 `json:"avg_start_time,omitempty"`
}
