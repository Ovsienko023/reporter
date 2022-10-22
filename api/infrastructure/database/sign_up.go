package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

const sqlSignUp = `
    insert into main.users (display_name, login, hash)
    values ($1, $2, $3)`

const sqlCheckLogin = `
   select 1 from main.users 
            where login = $1`

// SignUp возвращает следующие ошибки:
// ErrInternal
// ErrLoginAlreadyInUse
func (c *Client) SignUp(ctx context.Context, msg *SignUp) error {
	rowCheck, err := c.driver.Query(ctx, sqlCheckLogin, msg.Login)
	if err != nil {
		return NewInternalError(err)
	}

	var isUseLogin *int

	for rowCheck.Next() {
		err := rowCheck.Scan(&isUseLogin)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if isUseLogin != nil {
		return ErrLoginAlreadyInUse
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return NewInternalError(err)
	}

	row, err := transaction.Query(ctx, sqlSignUp,
		msg.DisplayName,
		msg.Login,
		msg.Password,
	)
	if err != nil {
		return NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Insert() {
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
