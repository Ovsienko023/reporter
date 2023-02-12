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
			   role_id,
			   created_at
from main.users as u
         left join main.users_to_roles as utr on u.id = utr.user_id

where u.deleted_at is null
  and (
        $3::uuid is null or
        u.id in (select pto.object_id
               from main.permissions_users_to_objects as pto
               where pto.object_type = 'users'
                 and pto.user_id = $3
                 and pto.object_id != pto.user_id)
    )
	)
	select (select count(*) from tab) as count,
			r.id                      as report_id,
       		r.display_name            as display_name,
			r.role_id                 as role, 
			r.created_at              as created_at
	from tab as r
	limit $1 offset $1 * ($2 - 1)`
)

func (c *Client) GetUsers(ctx context.Context, msg *GetUsers) ([]UserItem, *int, error) {
	row, err := c.driver.Query(ctx, sqlGetUsers,
		msg.PageSize,
		msg.Page,
		msg.AllowedTo,
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
			&user.Role,
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
	InvokerId string  `json:"invoker_id,omitempty"`
	Page      *int    `json:"page,omitempty"`
	PageSize  *int    `json:"page_size,omitempty"`
	AllowedTo *string `json:"allowed_to,omitempty"`
}

type UserItem struct {
	Id          *string    `json:"id,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	Role        *string    `json:"role,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
