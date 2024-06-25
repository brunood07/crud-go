package usecase

import (
	"crud/models"
	"crud/repositories"
)

type notificationsUsecase struct {
	repository repositories.NotificationsRepository
}

func NewNotificationsUsecase(repo repositories.NotificationsRepository) NotificationsUsecase {
	return &notificationsUsecase{
		repository: repo,
	}
}

func (nu *notificationsUsecase) CreateNotification(newNotification models.Notification) (models.Notification, error) {
	notification, err := nu.repository.CreateNotification(newNotification)
	if err != nil {
		return models.Notification{}, err
	}

	return notification, nil
}

func (nu *notificationsUsecase) SetNotificationRead(notificationId int) (models.Notification, error) {
	notification, err := nu.repository.SetNotificationReadByID(notificationId)
	if err != nil {
		return models.Notification{}, err
	}

	return notification, nil
}

func (nu *notificationsUsecase) GetUserNotifications(userId int) ([]models.Notification, error) {
	notifications, err := nu.repository.GetAllNotificationsForRecipient(userId)
	if err != nil {
		return []models.Notification{}, err
	}

	return notifications, nil
}

func (nu *notificationsUsecase) GetUserUnreadNotifications(userId int) ([]models.Notification, error) {
	notifications, err := nu.repository.GetRecipientUnreadNotifications(userId)
	if err != nil {
		return []models.Notification{}, err
	}

	return notifications, nil
}