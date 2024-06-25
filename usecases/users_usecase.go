package usecase

import (
	"crud/models"
	"crud/repositories"
)

type UsersUsecase struct {
	repository repositories.UsersRepository
}

func NewUsersUsecase(repo repositories.UsersRepository) UsersUsecase {
	return UsersUsecase{
		repository: repo,
	}
}

func (uu *UsersUsecase) CreateUser(newUser models.User) (models.User, error) {
	user, err := uu.repository.CreateUser(newUser)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (uu *UsersUsecase) GetUsers() ([]models.User, error) {
		users, err := uu.repository.GetUsers()
		if err != nil {
			return []models.User{}, err
		}

		return users, nil
}

func (uu *UsersUsecase) UpdateUser(id int, updateUser models.User) (models.User, error) {
	updatedUser, err := uu.repository.UpdateUser(id, updateUser)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (uu *UsersUsecase) DeleteUser(id int) (string, error) {
	delete, err := uu.repository.DeleteUser(id)
	if err != nil {
		return "", err
	}

	return delete, nil
}