package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type ShopService struct {
	db database.Shops
}

func NewShopService(db database.Shops) *ShopService {
	return &ShopService{db: db}
}

func (s *ShopService) AddNewShop(params *models.Shop) (*models.Shop, error) {
	return s.db.AddNewShop(params)
}

func (s *ShopService) UpdateShopParams(params *models.Shop) error {
	return s.db.UpdateShopParams(params)
}

func (s *ShopService) GetAllShops(page, limit int, search string) ([]*models.Shop, int, error) {
	offset := (page * limit) - limit
	return s.db.GetAllShops(limit, offset, search)
}
func (s *ShopService) GetShop(shopId int) (*models.Shop, error) {
	return s.db.GetShop(shopId)
}
func (s *ShopService) DeleteShop(shopId int) error {
	return s.db.DeleteShop(shopId)
}
