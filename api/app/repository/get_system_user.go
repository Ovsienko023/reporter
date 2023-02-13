package repository

import (
	"context"
)

//func (c *Client) GetSystemUser(ctx context.Context) (*SystemUser, error) {
//	return &SystemUser{
//		UserId:      ptr.String("08bc6135-dcb1-4ebe-bc3a-cfaa88db138f"), // todo
//		DisplayName: ptr.String("SystemUser"),
//	}
//}

const sqlGetSystem = `
	select id,
	       display_name
	from main.users
	where display_name = 'Administrator' and 
	      id = creator_id`

// GetSystemUser возвращает следующие ошибки:
//
// database.ErrInternal
// database.ErrUnexpectedBehavior
// database.ErrUserNotFound
func (c *Client) GetSystemUser(ctx context.Context) (*SystemUser, error) {
	row, err := c.driver.Query(ctx, sqlGetSystem)
	if err != nil {
		return nil, NewInternalError(err)
	}

	user := &SystemUser{}

	for row.Next() {
		err := row.Scan(
			&user.UserId,
			&user.DisplayName,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if user.UserId == nil {
		return nil, ErrInternal // todo
	}

	return user, nil
}

type SystemUser struct {
	UserId      *string `json:"user_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}
