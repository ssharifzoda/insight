package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Notification struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
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

type NotificationInfo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Shops       *ShopList `json:"shops"`
}

type ShopInfo struct {
	ShopId   int    `json:"shop_id"`
	ShopName string `json:"shop_name"`
}

type ShopList []*ShopInfo

func (sl *ShopList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to []byte", value)
	}
	return json.Unmarshal(bytes, sl)
}

func (sl *ShopList) Value() (driver.Value, error) {
	return json.Marshal(sl)
}
