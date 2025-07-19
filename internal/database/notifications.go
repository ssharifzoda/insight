package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
)

type NotificationDb struct {
	conn *gorm.DB
}

func NewNotificationDb(conn *gorm.DB) *NotificationDb {
	return &NotificationDb{conn: conn}
}

func (n *NotificationDb) CreateNewNotification(message *models.NotificationInput) error {
	var (
		notify         models.Notification
		notifyReceiver []*models.NotificationShop
	)
	notify.Image = message.Image
	notify.Description = message.Description
	notify.Title = message.Title
	tx := n.conn.Begin()
	err := tx.Create(&notify).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for i := range message.Shops {
		item := &models.NotificationShop{NotificationId: notify.Id, ShopId: i}
		notifyReceiver = append(notifyReceiver, item)
	}
	err = tx.Create(&notifyReceiver).Error
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
