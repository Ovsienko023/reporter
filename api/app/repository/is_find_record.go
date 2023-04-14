package repository

import (
	"context"
)

const sqlFindReport = `
	select 1
    from main.reports
	where id = $1 and creator_id = $2`

const sqlFindDayOff = `
	select 1
    from main.day_off
	where id = $1 and creator_id = $2`

const sqlFindSickLeave = `
	select 1
    from main.sick_leave
	where id = $1 and creator_id = $2`

const sqlFindVacationPaid = `
	select 1
    from main.vacations_paid
	where id = $1 and creator_id = $2`

const sqlFindVacationUnpaid = `
	select 1
    from main.vacations_unpaid
	where id = $1 and creator_id = $2`

func (c *Client) isFindReport(ctx context.Context, invokerId, reportId string) error {
	row, err := c.driver.Query(ctx, sqlFindReport, reportId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrReportIdNotFound
	}

	return nil
}

func (c *Client) isFindDayOff(ctx context.Context, invokerId, dayOffId string) error {
	row, err := c.driver.Query(ctx, sqlFindDayOff, dayOffId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrSickLeaveIdNotFound
	}

	return nil
}

func (c *Client) isFindSickLeave(ctx context.Context, invokerId, sickLeaveId string) error {
	row, err := c.driver.Query(ctx, sqlFindSickLeave, sickLeaveId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrSickLeaveIdNotFound
	}

	return nil
}

func (c *Client) isFindVacationPaid(ctx context.Context, invokerId, vacationPaidId string) error {
	row, err := c.driver.Query(ctx, sqlFindVacationPaid, vacationPaidId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrVacationIdNotFound
	}

	return nil
}

func (c *Client) isFindVacationUnpaid(ctx context.Context, invokerId, vacationUnpaidId string) error {
	row, err := c.driver.Query(ctx, sqlFindVacationUnpaid, vacationUnpaidId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrVacationIdNotFound
	}

	return nil
}
