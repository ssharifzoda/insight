package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"time"
)

type NotificationService struct {
	db database.Notifications
}

func NewNotificationService(db database.Notifications) *NotificationService {
	return &NotificationService{db: db}
}

func (n *NotificationService) CreateNewNotification(message *models.NotificationInput) error {
	path := consts.GlobalNotifyImagePath
	fileName := consts.NotificationPrefix + time.Now().Format("2006-01-02T15-04-05-0700")
	err := utils.SaveImageFromBase64(message.Image, path+fileName)
	if err != nil {
		return err
	}
	message.Image = path + fileName
	err = n.db.CreateNewNotification(message)
	if err != nil {
		utils.RemoveFile(path, fileName)
		return err
	}
	return nil
}
func (n *NotificationService) GetAllNotifications(page, limit int) ([]*models.Notification, error) {
	return n.db.GetAllNotifications(limit, page*limit-limit)
}

func (n *NotificationService) GetNotificationById(notificationId int) (*models.NotificationInfo, error) {
	return n.db.GetNotificationById(notificationId)
}

func (n *NotificationService) DeleteNotification(notificationId int) error {
	return n.db.DeleteNotification(notificationId)
}
