package repository

import (
	"context"
	"time"
)

// todo Добавить проверку на администратора
const (
	sqlGetUsers = `
	with tab as (
		select id,
			   display_name,
			   login,
			   created_at
		from main.users as u
		where u.deleted_at is null
	)
	select (select count(*) from tab) as count,
			r.id                      as report_id,
       		r.display_name            as display_name,
			r.login                   as login,
			r.created_at              as created_at
	from tab as r
	limit $1 offset $1 * ($2 - 1)`
)

func (c *Client) GetUsers(ctx context.Context, msg *GetUsers) ([]UserItem, *int, error) {
	row, err := c.driver.Query(ctx, sqlGetUsers,
		msg.PageSize,
		msg.Page,
	)
	if err != nil {
		return nil, nil, NewInternalError(err)
	}

	var count int
	users := make([]UserItem, 0, 0)

	for row.Next() {
		user := UserItem{}
		err := row.Scan(
			&count,
			&user.Id,
			&user.DisplayName,
			&user.Login,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, nil, NewInternalError(err)
		}
		users = append(users, user)
	}

	return users, &count, nil
}

type GetUsers struct {
	InvokerId string `json:"invoker_id,omitempty"`
	Page      *int   `json:"page,omitempty"`
	PageSize  *int   `json:"page_size,omitempty"`
}

type UserItem struct {
	Id          *string    `json:"id,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	Login       *string    `json:"login,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
