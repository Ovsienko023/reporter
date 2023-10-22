package postgres

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/infrastructure/database/dberrors"
	"github.com/Ovsienko023/reporter/infrastructure/database/dbmessage"
)

const sqlCreateUser = `
	select error
	from main.create_user(
	    _invoker_id := $1, 
	    _login := $2, 
	    _password := $3, 
	    _display_name := $4
	);`

func (s *Store) CreateUser(ctx context.Context, msg dbmessage.CreateUser) (*dbmessage.CreatedUser, error) {
	raw, err := s.driver.Query(ctx, sqlCreateUser,
		msg.InvokerId,
		msg.DisplayName,
	)
	if err != nil {
		return nil, dberrors.ErrInternal
	}

	if !raw.Next() {
		if err = raw.Err(); err != nil {
			return nil, dberrors.ErrInternal
		}
	}

	var (
		queryErr []byte
		result   dbmessage.CreatedUser
	)

	err = raw.Scan(
		&queryErr,
		&result.UserId,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", dberrors.ErrInternal, err)
	}

	if queryErr != nil {
		return nil, dberrors.ErrInternal
	}

	return &result, nil
}
