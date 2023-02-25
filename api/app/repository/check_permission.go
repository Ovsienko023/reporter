package repository

import "context"

const sqlCheckPermission = `
	select 1
    from main.permissions_users_to_objects
    where object_type = 'users' and
          user_id = $1 and
          object_id = $2`

func (c *Client) checkUserPermission(ctx context.Context, invokerId string, userId string) (bool, error) {
	if invokerId == userId { // todo fix permission
		return true, nil
	}

	raw, err := c.driver.Query(ctx, sqlCheckPermission, invokerId, userId)
	if err != nil {
		return false, NewInternalError(err)
	}

	raw.Next()
	status := raw.CommandTag()

	if status != nil && status.String() == "SELECT 1" {
		return true, nil
	}

	return false, nil
}
