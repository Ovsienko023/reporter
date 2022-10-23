package database

import "context"

const sqlGetProfile = `
	select id,
	       display_name,
	       login
    from main.users
    where id = $1`

func (c *Client) GetProfile(ctx context.Context, msg *GetProfile) (*Profile, error) {
	row, err := c.driver.Query(ctx, sqlGetProfile, msg.InvokerId)
	if err != nil {
		return nil, NewInternalError(err)
	}

	profile := &Profile{}

	for row.Next() {
		err := row.Scan(
			&profile.Id,
			&profile.DisplayName,
			&profile.Login,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	return profile, nil
}

type GetProfile struct {
	InvokerId string `yaml:"invoker_id,omitempty"`
}

type Profile struct {
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Login       *string `json:"login,omitempty"`
}
