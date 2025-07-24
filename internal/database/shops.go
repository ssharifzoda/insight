package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"time"
)

type ShopDb struct {
	conn *gorm.DB
}

func NewShopDb(conn *gorm.DB) *ShopDb {
	return &ShopDb{conn: conn}
}

func (s *ShopDb) AddNewShop(params *models.Shop) (*models.Shop, error) {
	return params, s.conn.Table("shops").Create(&params).Error
}

func (s *ShopDb) UpdateShopParams(params *models.Shop) error {
	return s.conn.Table("shops").Where("id").Updates(*params).Error
}

func (s *ShopDb) GetAllShops(limit, offset int, search string) ([]*models.Shop, int, error) {
	var shops []*models.Shop
	tx := s.conn.Table("shops").Where("status", 1)
	if search != "" {
		tx = tx.Where("fullname LIKE ?", "%"+search+"%").Find(&shops)
		return shops, 0, tx.Error
	}
	tx = tx.Limit(limit).Offset(offset).Find(&shops)
	var count int64
	tx.Count(&count)
	return shops, int(count), tx.Error
}
func (s *ShopDb) GetShop(shopId int) (*models.Shop, error) {
	var shop *models.Shop
	err := s.conn.Table("shops").Where("active = 1 and id = ?", shopId).First(&shop).Error
	return shop, err
}

func (s *ShopDb) DeleteShop(shopId int) error {
	return s.conn.Table("shops").Where("id", shopId).UpdateColumns(map[string]interface{}{
		"active":     0,
		"deleted_at": time.Now(),
	}).Error
}
