package repository

import "context"

const sqlCheckRole = `
	select role_id
    from main.users_to_roles
    where user_id = $1`

func (c *Client) checkAdminRole(ctx context.Context, invokerId string) (bool, error) {
	row, err := c.driver.Query(ctx, sqlCheckRole, invokerId)
	if err != nil {
		return false, NewInternalError(err)
	}

	var roleId *string

	for row.Next() {
		err := row.Scan(
			&roleId,
		)
		if err != nil {
			return false, NewInternalError(err)
		}
	}

	if *roleId != "administrator" {
		return false, nil
	}

	return true, nil
}
