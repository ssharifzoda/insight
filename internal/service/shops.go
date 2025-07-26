package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
)

type ShopService struct {
	db database.Shops
}

func NewShopService(db database.Shops) *ShopService {
	return &ShopService{db: db}
}

func (s *ShopService) AddNewShop(params *models.Shop) (*models.Shop, error) {
	if params.CompanyLogo != "" {
		filename := utils.FilePathGen("")
		path := consts.GlobalFilePath
		err := utils.SaveImageFromBase64(params.CompanyLogo, path+filename)
		if err != nil {
			return nil, err
		}
		params.CompanyLogo = filename
		return s.db.AddNewShop(params)
	}
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
	shop, err := s.db.GetShop(shopId)
	if err != nil {
		return nil, err
	}
	if shop.CompanyLogo != "" {
		file, err := utils.ConvertImageToBase64(consts.GlobalFilePath, shop.CompanyLogo)
		if err != nil {
			return nil, err
		}
		shop.CompanyLogo = file
		return shop, nil
	}
	return s.db.GetShop(shopId)
}
func (s *ShopService) DeleteShop(shopId int) error {
	return s.db.DeleteShop(shopId)
}
