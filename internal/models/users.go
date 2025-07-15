package models

import "time"

type User struct {
	Id        int        `json:"id"`
	FullName  string     `json:"full_name"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	RoleId    int        `json:"role_id"`
	Position  string     `json:"position"`
	Password  string     `json:"password"`
	Active    uint       `json:"active"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
