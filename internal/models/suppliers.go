package models

import (
	"time"
)

type Supplier struct {
	Id           int       `json:"id"`
	OrderNumber  int       `json:"order_number"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	CityId       int       `json:"city_id"`
	About        string    `json:"about"`
	Email        string    `json:"email"`
	TIN          string    `json:"TIN"`
	CompanyName  string    `json:"company_name"`
	CompanyLogo  string    `json:"company_logo"`
	LegalEntity  string    `json:"legal_entity"`
	MinimumOrder int       `json:"minimum_order"`
	Margin       int       `json:"margin"`
	SupplyDates  int       `json:"supply_dates"`
	Status       int       `json:"status"`
	VacationAt   time.Time `json:"vacation_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
