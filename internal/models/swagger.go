package models

import "github.com/lib/pq"

type UserInput struct {
	Id       int           `json:"id"`
	FullName string        `json:"full_name"`
	City     int           `json:"city"`
	Address  string        `json:"address"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Role     int           `json:"role"`
	Gender   string        `json:"gender"`
	Crops    pq.Int64Array `json:"crops" gorm:"type:int"`
	Username string        `json:"username"`
	Password string        `json:"password"`
}

type UserInputSW struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	City     int    `json:"city"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
	Gender   string `json:"gender"`
	Crops    []int  `json:"crops" gorm:"type:int"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReportInput struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Crops       int    `json:"crops"`
}

type ChangePasswordSW struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ReportUpdate struct {
	Id             int    `json:"id"`
	Image          string `json:"image"`
	Description    string `json:"description"`
	Crops          int    `json:"crops"`
	PlantDiseases  int    `json:"plant_diseases"`
	DisDescription string `json:"dis_description"`
}
