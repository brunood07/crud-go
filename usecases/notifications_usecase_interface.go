package usecase

import "crud/models"

type NotificationsUsecase interface {
	CreateNotification(newNotification models.Notification) (models.Notification, error)
	SetNotificationRead(notificationId int) (models.Notification, error)
	GetUserNotifications(userId int) ([]models.Notification, error)
}