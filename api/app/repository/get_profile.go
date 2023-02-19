package repository

import (
	"context"
	"time"
)

const sqlGetProfile = `
	select u.id,
		   u.display_name,
		   u.creator_id,
		   u.created_at,
		   login,
		   u.payload
    from main.users u
    inner join main.user_passwords on u.id = user_passwords.user_id 
        inner join main.user_logins on user_passwords.id = user_logins.grant_id
    where u.id = $1 and
          u.deleted_at is null`

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
			&profile.CreatorId,
			&profile.CreatedAt,
			&profile.Login,
			&profile.Payload,
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
	Id          *string      `json:"id,omitempty"`
	DisplayName *string      `json:"display_name,omitempty"`
	CreatorId   *string      `json:"creator_id,omitempty"`
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	Login       *string      `json:"login,omitempty"`
	Payload     *UserPayload `json:"payload,omitempty"`
}

type UserPayload struct{}
