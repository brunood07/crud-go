package usecase

import "crud/models"

type NotificationsUsecase interface {
	CreateNotification(newNotification models.Notification) (models.Notification, error)
}