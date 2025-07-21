package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"insight/pkg/consts"
	"time"
)

type NotificationDb struct {
	conn *gorm.DB
}

func NewNotificationDb(conn *gorm.DB) *NotificationDb {
	return &NotificationDb{conn: conn}
}

func (n *NotificationDb) CreateNewNotification(message *models.NotificationInput) (int, error) {
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
		return 0, err
	}
	for _, shop := range message.Shops {
		item := &models.NotificationShop{NotificationId: notify.Id, ShopId: shop}
		notifyReceiver = append(notifyReceiver, item)
	}
	err = tx.Table("notification_shop").Create(&notifyReceiver).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return notify.Id, nil
}

func (n *NotificationDb) GetAllNotifications(limit, offset int) ([]*models.Notification, error) {
	var result []*models.Notification
	err := n.conn.Where("status = 1").Limit(limit).Offset(offset).Find(&result).Error
	return result, err
}

func (n *NotificationDb) GetNotificationById(notificationId int) (*models.NotificationInfo, error) {
	var info *models.NotificationInfo
	err := n.conn.Raw(consts.GetNotificationInfoSQL, notificationId).Scan(&info).Error
	return info, err
}
func (n *NotificationDb) DeleteNotification(notificationId int) error {
	err := n.conn.Table("notifications").Where("id", notificationId).UpdateColumns(map[string]interface{}{
		"status":     0,
		"updated_at": time.Now(),
	}).Error
	return err
}
