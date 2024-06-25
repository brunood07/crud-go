package models

import "time"

type Notification struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
	RecipientId int `json:"recipient_id"`
}

type CreateNotification struct {
	Title string `json:"title"`
	Content string `json:"content"`
}