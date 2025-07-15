package models

import "time"

type SingIn struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ChangePassword struct {
	UserId      int    `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserAuth struct {
	UserId        int        `json:"user_id"`
	AccessToken   string     `json:"access_token"`
	RefreshToken  string     `json:"refresh_token"`
	TemporaryPass int        `json:"temporary_pass"`
	PassResetAt   *time.Time `json:"pass_reset_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
