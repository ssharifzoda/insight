package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"time"
)

type ProductService struct {
	db database.Products
}

func NewProductService(db database.Products) *ProductService {
	return &ProductService{db: db}
}

func (p *ProductService) AddNewProduct(product *models.Product) error {
	path := consts.GlobalLogoFilePath + product.Name + time.Now().Format(time.DateOnly)
	err := utils.SaveImageFromBase64(product.Image, path)
	if err != nil {
		return err
	}
	product.Image = path
	return p.db.AddNewProduct(product)
}

func (p *ProductService) GetAllProducts(page, limit int, filter *models.ProductFilter) ([]*models.Product, error) {
	offset := (page * limit) - limit
	return p.db.GetAllProducts(limit, offset, filter)
}
func (p *ProductService) GetProductById(productId int) (*models.Product, error) {
	product, err := p.db.GetProductById(productId)
	if err != nil {
		return nil, err
	}
	product.Image, err = utils.ConvertImageToBase64(consts.GlobalLogoFilePath, product.Image)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) EditProduct(product *models.Product) error {
	item, err := p.db.GetProductById(product.Id)
	if err != nil {
		return err
	}
	err = utils.RemoveFile(consts.GlobalLogoFilePath, item.Image)
	if err != nil {
		return err
	}
	path := consts.GlobalLogoFilePath + product.Name + time.Now().Format(time.DateOnly)
	err = utils.SaveImageFromBase64(product.Image, path)
	if err != nil {
		return err
	}
	product.Image = path
	return p.db.EditProduct(product)
}

func (p *ProductService) DeleteProduct(productId int) error {
	item, err := p.db.GetProductById(productId)
	if err != nil {
		return err
	}
	err = utils.RemoveFile(consts.GlobalLogoFilePath, item.Image)
	if err != nil {
		return err
	}
	return p.db.DeleteProduct(productId)
}
