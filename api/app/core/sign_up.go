package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

//SignUp возвращает следующие ошибки:
//	ErrInternal
//	ErrLoginAlreadyInUse
func (c *Core) SignUp(ctx context.Context, msg *domain.SignUpRequest) error {
	err := c.db.SignUp(ctx, msg.ToDbSignUp())
	if err != nil {
		switch {
		case errors.Is(err, database.ErrLoginAlreadyInUse):
			return ErrLoginAlreadyInUse
		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}
