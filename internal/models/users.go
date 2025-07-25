package models

import "time"

type User struct {
	Id         int        `json:"id"`
	FullName   string     `json:"full_name"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	RoleId     int        `json:"role_id"`
	Position   string     `json:"position"`
	ShopId     int        `json:"shop_id,omitempty"`
	SupplierId int        `json:"supplier_id,omitempty"`
	Password   string     `json:"password"`
	Active     uint       `json:"active"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type UserInfo struct {
	Id         int        `json:"id"`
	FullName   string     `json:"full_name"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	RoleId     int        `json:"role_id"`
	Position   string     `json:"position"`
	Password   string     `json:"password"`
	ShopId     int        `json:"shop_id,omitempty"`
	SupplierId int        `json:"supplier_id,omitempty"`
	Active     uint       `json:"active"`
	Shop       *Shop      `json:"shop,omitempty" gorm:"-"`
	SalePoint  *SalePoint `json:"sale_point,omitempty" gorm:"-"`
	Supplier   *Supplier  `json:"supplier,omitempty" gorm:"-"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
