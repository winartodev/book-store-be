package entity

import "time"

type Publisher struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
