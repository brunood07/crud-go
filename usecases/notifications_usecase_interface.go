package usecase

import "crud/models"

type NotificationsUsecase interface {
	CreateNotification(newNotification models.CreateNotification) (models.Notification, error)
	SetNotificationRead(notificationId int) (models.Notification, error)
	GetUserNotifications(userId int) ([]models.Notification, error)
	GetUserUnreadNotifications(userId int) ([]models.Notification, error)
}