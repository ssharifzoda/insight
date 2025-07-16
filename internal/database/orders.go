package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"insight/pkg/consts"
)

type OrderDb struct {
	conn *gorm.DB
}

func NewOrderDb(conn *gorm.DB) *OrderDb {
	return &OrderDb{conn: conn}
}

func (o *OrderDb) AddNewOrder(orderParams *models.OrderInput) error {
	order := &models.Order{
		ShopId:      orderParams.ShopId,
		Total:       orderParams.Total,
		Comments:    orderParams.Comments,
		SupComments: orderParams.SupComments,
		Canceled:    orderParams.Canceled,
		WhoCanceled: orderParams.WhoCanceled,
	}
	tx := o.conn.Begin()
	err := tx.Table("orders").Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var items []*models.OrderProducts
	for _, item := range orderParams.Products {
		orderProducts := &models.OrderProducts{
			OrderId:   order.Id,
			ProductId: item.ProductId,
			Qty:       item.Qty,
			Price:     item.Price,
			Status:    1,
		}
		items = append(items, orderProducts)
	}
	err = tx.Table("order_products").Create(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (o *OrderDb) GetAllOrders(filter *models.OrderFilter) (orders []*models.Order, err error) {
	tx := o.conn.Table("orders as o")
	if filter.Status != nil {
		tx = tx.Where("o.status", filter.Status)
	}
	if filter.ShopId != 0 {
		tx = tx.Where("o.shop_id", filter.ShopId)
	}
	if filter.SupplierId != 0 {
		tx.InnerJoins("JOIN order_products as op on op.order_id = o.id")
		tx = tx.InnerJoins("JOIN products as p on op.product_id = p.id")
		tx = tx.Where("p.supplier_id", filter.SupplierId)
	}
	if filter.DateFrom != "" { //todo: in service part validate
		tx = tx.Where("o.created_at between ? and ?", filter.DateFrom, filter.DateTo)
	}
	tx = tx.Find(&orders)
	return orders, nil
}

func (o *OrderDb) GetOrderById(orderId int) (order *models.OrderInfo, err error) {
	err = o.conn.Raw(consts.GetOrderByIdSQL, orderId).First(&order).Error
	return order, err
}

func (o *OrderDb) EditOrder(order *models.OrderInput) error {
	orderParams := &models.Order{
		Id:          order.Id,
		ShopId:      order.ShopId,
		Total:       order.Total,
		Comments:    order.Comments,
		SupComments: order.SupComments,
		Canceled:    order.Canceled,
		WhoCanceled: order.WhoCanceled,
		Status:      order.Status,
	}
	tx := o.conn.Begin()
	err := tx.Table("orders").Updates(&orderParams).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var items []*models.OrderProducts
	for _, item := range order.Products {
		orderProducts := &models.OrderProducts{
			OrderId:   order.Id,
			ProductId: item.ProductId,
			Qty:       item.Qty,
			Price:     item.Price,
			Status:    1,
		}
		items = append(items, orderProducts)
	}

	err = tx.Where("order_id = ?", order.Id).Delete(&models.OrderProducts{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("order_products").Save(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
