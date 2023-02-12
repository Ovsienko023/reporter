package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
)

const sqlSignUp = `
    insert into main.users (display_name, login, hash)
    values ($1, $2, $3)
    returning id`

const sqlCheckLogin = `
   select id
   from main.users 
	where login = $1
	limit 1`

const sqlAddDefRole = `
	insert into main.users_to_roles (user_id, role_id)
	values ($1, 'default')`

const sqlAddPermission = `
	insert into main.permissions_users_to_objects (user_id, object_type, object_id)
	values ($1, 'users', $1)`

// SignUp возвращает следующие ошибки:
// ErrInternal
// ErrLoginAlreadyInUse
func (c *Client) SignUp(ctx context.Context, msg *SignUp) error {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return NewInternalError(err)
	}

	defer transaction.Rollback(ctx)

	rawCheck, err := transaction.Exec(ctx, sqlCheckLogin, msg.Login)
	if err != nil {
		return NewInternalError(err)
	}

	if rawCheck.String() == "SELECT 1" {
		return ErrLoginAlreadyInUse
	}

	rawCreate, err := transaction.Query(ctx, sqlSignUp,
		msg.DisplayName,
		msg.Login,
		msg.Password,
	)
	if err != nil {
		return NewInternalError(err)
	}

	var userId *string

	for rawCreate.Next() {
		err := rawCreate.Scan(&userId)
		if err != nil {
			return NewInternalError(err)
		}
	}

	rawAddRole, err := transaction.Exec(ctx, sqlAddDefRole, userId)
	if !rawAddRole.Insert() {
		return NewInternalError(err)
	}

	rawAdd, err := transaction.Exec(ctx, sqlAddPermission, userId)
	if !rawAdd.Insert() {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)
	return nil
}

type SignUp struct {
	Login       string  `json:"login,omitempty"`
	Password    string  `json:"password,omitempty"`
	DisplayName *string `json:"display_name,"`
}
