package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type OrderService struct {
	db database.Orders
}

func NewOrderService(db database.Orders) *OrderService {
	return &OrderService{db: db}
}

func (o *OrderService) AddNewOrder(order *models.OrderInput) error {
	return o.db.AddNewOrder(order)
}
func (o *OrderService) GetAllOrders(filter *models.OrderFilter) (orders []*models.Order, err error) {
	return o.db.GetAllOrders(filter)
}

func (o *OrderService) GetOrderById(orderId int) (order *models.OrderInfo, err error) {

	order, err = o.db.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	for _, product := range order.Products {
		product.Total = product.Price * float64(product.Qty)
	}
	return order, err
}

func (o *OrderService) EditOrder(order *models.OrderInput) error {
	return o.db.EditOrder(order)
}
