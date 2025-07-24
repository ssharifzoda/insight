package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"time"
)

type ProductDb struct {
	conn *gorm.DB
}

func NewProductDb(conn *gorm.DB) *ProductDb {
	return &ProductDb{conn: conn}
}

func (p *ProductDb) AddNewProduct(product *models.Product) (*models.Product, error) {
	return product, p.conn.Table("products").Create(&product).Error
}

func (p *ProductDb) GetAllProducts(limit, offset int, filter *models.ProductFilter) (result []*models.Product, totalCount int, err error) {
	tx := p.conn.Table("products").Where("status = 1")
	if filter != nil {
		if filter.Search != "" {
			tx = tx.Where("name LIKE ?", "%"+filter.Search+"%").Find(&result)
			return result, 0, tx.Error
		}
		if filter.SupplierId != 0 {
			tx = tx.Where("supplier_id = ?", filter.SupplierId)
		}
		if filter.BrandId != 0 {
			tx = tx.Where("brand_id = ?", filter.BrandId)
		}
		if filter.CategoryId != 0 {
			tx = tx.Where("category_id = ?", filter.CategoryId)
		}
	}
	tx = tx.Limit(limit).Offset(offset).Order("order_number DESC").Find(&result)
	var count int64
	tx.Count(&count)
	return result, int(count), tx.Error
}
func (p *ProductDb) GetProductById(productId int) (result *models.Product, err error) {
	return result, p.conn.Table("products").Where("status = 1 and id = ?", productId).First(&result).Error
}

func (p *ProductDb) EditProduct(product *models.Product) error {
	return p.conn.Table("products").Updates(&product).Error
}

func (p *ProductDb) DeleteProduct(productId int) error {
	return p.conn.Table("products").Where("id", productId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"deleted_at": time.Now(),
	}).Error
}
