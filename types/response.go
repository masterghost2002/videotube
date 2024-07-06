package types

import "database/sql"

type UserResponse struct {
	ID         int64          `json:"id"`
	Username   string         `json:"username"`
	FullName   string         `json:"full_name"`
	Email      string         `json:"email"`
	Profileurl sql.NullString `json:"profile_url"`
	ChannelID  sql.NullInt64  `json:"channel_id"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}
