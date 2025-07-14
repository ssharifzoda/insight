package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type ProductService struct {
	db database.Products
}

func NewProductService(db database.Products) *ProductService {
	return &ProductService{db: db}
}

func (p *ProductService) AddNewProduct(product *models.Product) error {
	return p.db.AddNewProduct(product)
}

func (p *ProductService) GetAllProducts(page, limit int) ([]*models.Product, error) {
	offset := (page * limit) - limit
	return p.db.GetAllProducts(limit, offset)
}
func (p *ProductService) GetProductById(productId int) (*models.Product, error) {
	return p.db.GetProductById(productId)
}

func (p *ProductService) EditProduct(product *models.Product) error {
	return p.db.EditProduct(product)
}

func (p *ProductService) DeleteProduct(productId int) error {
	return p.db.DeleteProduct(productId)
}
