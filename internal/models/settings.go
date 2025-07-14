package models

import (
	"time"
)

type Brand struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Logo      string     `json:"logo"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

type Category struct {
	Id          int        `json:"id"`
	ParentId    int        `json:"parent_id"`
	OrderNumber int        `json:"order_number"`
	Name        string     `json:"name"`
	Logo        string     `json:"logo"`
	Status      int        `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   time.Time  `json:"deleted_at"`
}

type City struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

type Promotion struct {
	Id         int        `json:"id"`
	SupplierId int        `json:"supplier_id"`
	Name       string     `json:"name"`
	FromDate   time.Time  `json:"from_date"`
	ToDate     time.Time  `json:"to_date"`
	BuyQty     int        `json:"buy_qty"`
	GiftQty    int        `json:"gift_qty"`
	Status     int        `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
}
type PromotionInfo struct {
	Id         int       `json:"id"`
	SupplierId int       `json:"supplier_id"`
	Name       string    `json:"name"`
	FromDate   time.Time `json:"from_date"`
	ToDate     time.Time `json:"to_date"`
	BuyQty     int       `json:"buy_qty"`
	GiftQty    int       `json:"gift_qty"`
	Products   []string  `json:"products"`
}

type Role struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

type RoleInput struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}

type RolePermission struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

type RoleInfo struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Status      int        `json:"status"`
	Permissions []string   `json:"permissions"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Permission struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}
