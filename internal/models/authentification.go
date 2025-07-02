package models

type SingIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	UserId      int    `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
