package entity

import "time"

type Book struct {
	ID          int64     `json:"id"`
	PublisherID int64     `json:"publisher_id"`
	CategoryID  int64     `json:"category_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publication int       `json:"year_of_publication"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
