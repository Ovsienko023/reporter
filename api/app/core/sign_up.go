package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

// SignUp возвращает следующие ошибки:
//
//	ErrInternal
//	ErrLoginAlreadyInUse
func (c *Core) SignUp(ctx context.Context, msg *domain.SignUpRequest) error {
	hash, err := c.generateHash(msg.Password)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	systemUser, _ := c.db.GetSystemUser(ctx)

	err = c.db.SignUp(ctx, msg.ToDbSignUp(*systemUser.UserId, hash))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrLoginAlreadyInUse):
			return ErrLoginAlreadyInUse
		case errors.Is(err, repository.ErrUnauthorized):
			return ErrUnauthorized

		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}
