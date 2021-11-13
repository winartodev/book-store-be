package bookstorebe

// Entity
type Publisher struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID          int64  `json:"id"`
	PublisherID int64  `json:"publisher_id"`
	CategoryID  int64  `json:"category_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication int    `json:"year_of_publication"`
	Stock       int    `json:"stock"`
}

// Repository

// Usecase
