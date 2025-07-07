package models

import "time"

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
