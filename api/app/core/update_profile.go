package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) UpdateProfile(ctx context.Context, msg *domain.UpdateProfileRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.UpdateProfile(ctx, msg.ToDbUpdateProfile(invokerId))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return nil
}
