package utils

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"insight/internal/models"
)

func FirebaseSender(info models.NotificationInfo, conn *messaging.Client) error {
	var devices []string
	for _, shop := range *info.Shops {
		devices = append(devices, shop.ShopName)
	}
	if len(devices) > 1 {
		message := &messaging.MulticastMessage{
			Tokens: devices,
			Notification: &messaging.Notification{
				Title: info.Title,
				Body:  info.Image,
			},
			Data: map[string]string{"type": "update"},
		}
		_, err := conn.SendMulticast(context.Background(), message)
		if err != nil {
			return err
		}
		return nil
	}
	message := &messaging.Message{
		Token: devices[0],
		Notification: &messaging.Notification{
			Title: info.Title,
			Body:  info.Description,
		},
		Data: map[string]string{"type": "update"},
	}
	_, err := conn.Send(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
