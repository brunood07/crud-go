package repositories

import (
	"crud/db"
	"crud/models"
	"database/sql"
	"fmt"
	"time"
)

type NotificationsRepository struct {
	connection *sql.DB
}

func NewNotificationsRepository(connection *sql.DB) NotificationsRepository {
	return NotificationsRepository{
		connection: connection,
	}
}

func (nr *NotificationsRepository) CreateNotification(newNotification models.CreateNotification) (models.Notification, error) {
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
		var notification models.Notification
		stmt, err := db.CON.Prepare("INSERT INTO notification (title, content, recipientId) VALUES ($1, $2, $3) RETURNING id")
		if err != nil {
			fmt.Println(err)
		return models.Notification{}, err
		}
		defer stmt.Close()

		err = stmt.QueryRow(notification.Title, notification.Content, user.ID).Scan(&notification.ID)
		if err != nil {
			fmt.Println(err)
		return models.Notification{}, err
		}
	}

	return models.Notification{}, nil
}

func (nr *NotificationsRepository) SetNotificationReadByID(id int) (models.Notification, error) {
	// Update the readAt field
	stmt, err := nr.connection.Prepare("UPDATE notification SET readAt=$1 WHERE id=$2")
	if err != nil {
		fmt.Println(err)
		return models.Notification{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		fmt.Println(err)
		return models.Notification{}, err
	}

	// Retrieve the updated notification
	var updatedNotification models.Notification
	query := "SELECT id, title, content, readAt, recipientId FROM notification WHERE id=$1"
	err = nr.connection.QueryRow(query, id).Scan(
		&updatedNotification.ID, &updatedNotification.Title, &updatedNotification.Content, &updatedNotification.ReadAt, &updatedNotification.RecipientId,
	)
	if err != nil {
		fmt.Println(err)
		return models.Notification{}, err
	}

	return updatedNotification, nil
}

func (nr *NotificationsRepository) GetAllNotificationsForRecipient(recipientId int) ([]models.Notification, error) {
	// Prepare the SELECT statement
	query := "SELECT id, title, content, readAt, recipientId FROM notification WHERE recipientId=$1"
	rows, err := nr.connection.Query(query, recipientId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.ID, &notification.Title, &notification.Content, &notification.ReadAt, &notification.RecipientId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return notifications, nil
}

func (nr *NotificationsRepository) GetRecipientUnreadNotifications(recipientId int) ([]models.Notification, error) {
		// Prepare the SELECT statement
		query := "SELECT id, title, content, readAt, recipientId FROM notification WHERE recipientId=$1 AND readAt is NULL"
		rows, err := nr.connection.Query(query, recipientId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer rows.Close()
	
		var notifications []models.Notification
		for rows.Next() {
			var notification models.Notification
			err := rows.Scan(&notification.ID, &notification.Title, &notification.Content, &notification.ReadAt, &notification.RecipientId)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			notifications = append(notifications, notification)
		}
	
		// Check for errors from iterating over rows
		if err = rows.Err(); err != nil {
			fmt.Println(err)
			return nil, err
		}
	
		return notifications, nil
}