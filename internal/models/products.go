package models

import "time"

type Product struct {
	Id          int        `json:"id"`
	OrderNumber int        `json:"order_number"`
	Name        string     `json:"name"`
	CategoryId  int        `json:"category_id"`
	BrandId     int        `json:"brand_id"`
	Description string     `json:"description"`
	PromotionId int        `json:"promotion_id"`
	Price       float64    `json:"price"`
	Image       string     `json:"image"`
	SupplierId  int        `json:"supplier_id"`
	Status      int        `json:"status"`
	Code        string     `json:"code"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
