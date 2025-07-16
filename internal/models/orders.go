package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	Id          int        `json:"id"`
	ShopId      int        `json:"shop_id"`
	Total       float64    `json:"total"`
	DeliverAt   *time.Time `json:"deliver_at"`
	Comments    string     `json:"comments"`
	SupComments string     `json:"sup_comments"`
	Status      int        `json:"status"`
	VerifiedAt  *time.Time `json:"verified_at"`
	DeliveredAt *time.Time `json:"delivered_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Canceled    string     `json:"canceled"`
	WhoCanceled string     `json:"who_canceled"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type OrderProducts struct {
	Id        int        `json:"id"`
	OrderId   int        `json:"order_id"`
	ProductId int        `json:"product_id"`
	Qty       int        `json:"qty"`
	Price     float64    `json:"price"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type OrderInfo struct {
	Id          int         `json:"id"`
	Shop        string      `json:"shop"`
	Total       float64     `json:"total"`
	DeliverAt   *time.Time  `json:"deliver_at"`
	Comments    string      `json:"comments"`
	SupComments string      `json:"sup_comments"`
	Status      int         `json:"status"`
	VerifiedAt  *time.Time  `json:"verified_at"`
	DeliveredAt *time.Time  `json:"delivered_at"`
	CompletedAt *time.Time  `json:"completed_at"`
	Canceled    string      `json:"canceled"`
	WhoCanceled string      `json:"who_canceled"`
	Products    ProductList `json:"products"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
}

type ProductList []*Products

func (pl *ProductList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to []byte", value)
	}
	return json.Unmarshal(bytes, pl)
}

func (pl *ProductList) Value() (driver.Value, error) {
	return json.Marshal(pl)
}

type OrderInput struct {
	Id          int     `json:"id"`
	ShopId      int     `json:"shop_id"`
	Total       float64 `json:"total"`
	Comments    string  `json:"comments"`
	SupComments string  `json:"sup_comments"`
	Canceled    string  `json:"canceled"`
	WhoCanceled string  `json:"who_canceled"`
	OrderId     int     `json:"order_id"`
	Status      int     `json:"status"`
	Products    []struct {
		ProductId int     `json:"product_id"`
		Qty       int     `json:"qty"`
		Price     float64 `json:"price"`
	} `json:"products"`
}

type Products struct {
	Product   string  `json:"product" `
	ProductId int     `json:"product_id"`
	Image     string  `json:"image"`
	Qty       int     `json:"qty"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
}

type OrderFilter struct {
	ShopId     int    `json:"shop_id"`
	SupplierId int    `json:"supplier_id"`
	Status     *int   `json:"status"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}
