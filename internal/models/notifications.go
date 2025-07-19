package models

import (
	"time"
)

type Notification struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	SupplierId  int        `json:"supplier_id"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type NotificationInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Shops       []int  `json:"shops"`
}

type NotificationShop struct {
	Id             int        `json:"id"`
	NotificationId int        `json:"notification_id"`
	ShopId         int        `json:"shop_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
