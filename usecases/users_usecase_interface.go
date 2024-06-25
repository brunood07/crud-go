package usecase

import "crud/models"

type UsersUsecase interface {
	CreateUser(newUser models.CreateUser) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id int, updateUser models.User) (models.User, error)
	DeleteUser(id int) (string, error)
}