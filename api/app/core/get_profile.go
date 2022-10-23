package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetProfile(ctx context.Context, msg *domain.GetProfileRequest) (*domain.GetProfileResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetProfile(ctx, msg.ToDbGetProfile(invokerId))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return domain.FromGetProfileResponse(result), nil
}
