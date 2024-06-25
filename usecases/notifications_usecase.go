package usecase

import (
	"crud/models"
	"crud/repositories"
)

type NotificationsUsecase struct {
	repository repositories.NotificationsRepository
}

func NewNotificationsUsecase(repo repositories.NotificationsRepository) NotificationsUsecase {
	return NotificationsUsecase{
		repository: repo,
	}
}

func (nu *NotificationsUsecase) CreateNotification(newNotification models.Notification) (models.Notification, error) {
	notification, err := nu.repository.CreateNotification(newNotification)
	if err != nil {
		return models.Notification{}, err
	}

	return notification, nil
}