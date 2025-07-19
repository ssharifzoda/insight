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
	fileName := consts.NotificationPrefix + time.Now().Format(time.RFC3339)
	err := utils.SaveImageFromBase64(message.Image, path+fileName)
	if err != nil {
		return err
	}
	message.Image = path
	err = n.db.CreateNewNotification(message)
	if err != nil {
		err = utils.RemoveFile(path, fileName)
		return err
	}
	return nil
}
