package repository

import "context"

const sqlCheckPermission = `
	select role_id
    from main.users_to_roles
    where user_id = $1`

func (c *Client) checkPermission(ctx context.Context, invokerId string) (bool, error) {
	row, err := c.driver.Query(ctx, sqlCheckPermission, invokerId)
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
