package usecase

import (
	"crud/models"
	"crud/repositories"
)

type usersUsecase struct {
	repository repositories.UsersRepository
}

func NewUsersUsecase(repo repositories.UsersRepository) UsersUsecase {
	return &usersUsecase{
		repository: repo,
	}
}

func (uu *usersUsecase) CreateUser(newUser models.CreateUser) (models.User, error) {
	user, err := uu.repository.CreateUser(newUser)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (uu *usersUsecase) GetUsers() ([]models.User, error) {
		users, err := uu.repository.GetUsers()
		if err != nil {
			return []models.User{}, err
		}

		return users, nil
}

func (uu *usersUsecase) UpdateUser(id int, updateUser models.User) (models.User, error) {
	updatedUser, err := uu.repository.UpdateUser(id, updateUser)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (uu *usersUsecase) DeleteUser(id int) (string, error) {
	delete, err := uu.repository.DeleteUser(id)
	if err != nil {
		return "", err
	}

	return delete, nil
}