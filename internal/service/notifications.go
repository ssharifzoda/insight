package service

import (
	"firebase.google.com/go/v4/messaging"
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
)

type NotificationService struct {
	db           database.Notifications
	firebaseConn *messaging.Client
}

func NewNotificationService(db database.Notifications, firebaseConn *messaging.Client) *NotificationService {
	return &NotificationService{db: db, firebaseConn: firebaseConn}
}

func (n *NotificationService) CreateNewNotification(message *models.NotificationInput) error {
	path := consts.GlobalFilePath
	fileName := utils.FilePathGen("notification")
	err := utils.SaveImageFromBase64(message.Image, path+fileName)
	if err != nil {
		return err
	}
	message.Image = fileName
	notifyId, err := n.db.CreateNewNotification(message)
	if err != nil {
		utils.RemoveFile(path, fileName)
		return err
	}
	notification, err := n.db.GetNotificationById(notifyId)
	if err != nil {
		utils.RemoveFile(path, fileName)
		return err
	}
	utils.FirebaseSender(*notification, n.firebaseConn)
	return nil
}
func (n *NotificationService) GetAllNotifications(page, limit int) ([]*models.Notification, int, error) {
	return n.db.GetAllNotifications(limit, page*limit-limit)
}

func (n *NotificationService) GetNotificationById(notificationId int) (*models.NotificationInfo, error) {
	notification, err := n.db.GetNotificationById(notificationId)
	if err != nil {
		return nil, err
	}
	notification.Image, err = utils.ConvertImageToBase64(consts.GlobalFilePath, notification.Image)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (n *NotificationService) DeleteNotification(notificationId int) error {
	return n.db.DeleteNotification(notificationId)
}
