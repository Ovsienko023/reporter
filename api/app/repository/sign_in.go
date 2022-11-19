package repository

import (
	"context"
)

const sqlSignIn = `
	select id,
	       hash 
    from main.users
    where login = $1 
    and deleted_at is null`

func (c *Client) SignIn(ctx context.Context, msg *SignIn) (*Auth, error) {
	row, err := c.driver.Query(ctx, sqlSignIn, msg.Login)
	if err != nil {
		return nil, NewInternalError(err)
	}

	auth := &Auth{}

	for row.Next() {
		err := row.Scan(
			&auth.UserId,
			&auth.Password,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if auth.Password == nil {
		return nil, ErrCredentials
	}

	return auth, nil
}

type SignIn struct {
	Login string `json:"login,omitempty"`
}

type Auth struct {
	UserId   *string `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}
