package models

import "time"

type Notification struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ReadAt time.Time `json:"readAt" valid:"null"`
	RecipientId int `json:"recipientId"`
}