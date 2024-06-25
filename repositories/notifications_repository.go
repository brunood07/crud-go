package repositories

import (
	"crud/db"
	"crud/models"
	"database/sql"
	"fmt"
)

type NotificationsRepository struct {
	connection *sql.DB
}

func NewNotificationsRepository(connection *sql.DB) NotificationsRepository {
	return NotificationsRepository{
		connection: connection,
	}
}

func (nr *NotificationsRepository) CreateNotification(newNotification models.Notification) (models.Notification, error) {

	rows, err := db.CON.Query("SELECT id, first_name, last_name, age, email FROM users")
	if err != nil {
		fmt.Println(err)
		return models.Notification{}, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Email)
		if err != nil {
			fmt.Println(err)
		return models.Notification{}, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return models.Notification{}, err
	}

	for _, user := range users {
		stmt, err := db.CON.Prepare("INSERT INTO notification (title, content, readAt, recipientId) VALUES ($1, $2, $3, $4) RETURNING id")
		if err != nil {
			fmt.Println(err)
		return models.Notification{}, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(newNotification.Title, newNotification.Content, newNotification.ReadAt, user.ID).Scan(&newNotification.ID)
		if err != nil {
			fmt.Println(err)
		return models.Notification{}, err
		}
	}

	return newNotification, nil
}