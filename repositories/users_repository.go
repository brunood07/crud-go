package repositories

import (
	"crud/db"
	"crud/models"
	"database/sql"
	"fmt"
)

type UsersRepository struct {
	connection *sql.DB
}

func NewUsersRepository(connection *sql.DB) UsersRepository {
	return UsersRepository{
		connection: connection,
	}
}

func (ur *UsersRepository) CreateUser(newUser models.User) (models.User, error) {
	
	stmt, err := ur.connection.Prepare("INSERT INTO users (first_name, last_name, age, email) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(newUser.FirstName, newUser.LastName, newUser.Age, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	stmt.Close()
	return newUser, nil
}

func (ur *UsersRepository) GetUsers() ([]models.User, error) {
	rows, err := ur.connection.Query("SELECT id, first_name, last_name, age, email FROM users")
	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Email)
		if err != nil {
			fmt.Println(err)
			return []models.User{}, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	return users, nil
}

func (ur *UsersRepository) UpdateUser(id int, updateUser models.User) (models.User, error) {
	stmt, err := ur.connection.Prepare("UPDATE users SET first_name=$1, last_name=$2, age=$3, email=$4 WHERE id=$5")
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateUser.FirstName, updateUser.LastName, updateUser.Age, updateUser.Email, id)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	var updatedUser models.User
	err = ur.connection.QueryRow("SELECT id, first_name, last_name, age, email FROM users WHERE id=$1", id).Scan(
		&updatedUser.ID, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Age, &updatedUser.Email,
	)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}

	return updatedUser, nil
}

func (ur *UsersRepository) DeleteUser(id int) (string, error) {
	stmt, err := db.CON.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return "user deleted", nil
}