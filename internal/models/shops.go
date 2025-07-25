package models

import "time"

type Shop struct {
	Id          int         `json:"id"`
	Fullname    string      `json:"fullname"`
	Address     string      `json:"address"`
	CityId      interface{} `json:"city_id" gorm:"type:int"`
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

type SalePoint struct {
	Id        int        `json:"id"`
	ShopId    int        `json:"shop_id"`
	Name      string     `json:"name"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
