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