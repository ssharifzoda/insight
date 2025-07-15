package models

import "time"

type Shop struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	CityId      interface{} `json:"city_id"`
	UserId      int         `json:"user_id"`
	About       string      `json:"about"`
	Email       string      `json:"email"`
	TIN         string      `json:"TIN"`
	CompanyName string      `json:"company_name"`
	CompanyLogo string      `json:"company_logo"`
	LegalEntity string      `json:"legal_entity"`
	Status      int         `json:"status"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
}
