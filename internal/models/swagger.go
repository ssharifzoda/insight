package models

type UserSW struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	Position string `json:"position"`
	Password string `json:"password"`
}

type ReportInput struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Crops       int    `json:"crops"`
}

type ReportUpdate struct {
	Id             int    `json:"id"`
	Image          string `json:"image"`
	Description    string `json:"description"`
	Crops          int    `json:"crops"`
	PlantDiseases  int    `json:"plant_diseases"`
	DisDescription string `json:"dis_description"`
}
