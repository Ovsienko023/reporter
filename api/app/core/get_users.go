package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetUsers(ctx context.Context, request *domain.GetUsersRequest) (*domain.GetUsersResponse, error) {
	invokerId, err := c.authorize(request.Token)
	if err != nil {
		return nil, err
	}

	result, cnt, err := c.db.GetUsers(ctx, request.ToDbGetUsers(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return domain.FromGetUsersResponse(result, cnt), nil
}
