package bookstorebe

import "context"

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
type PublisherRepository interface {
	// seller
	GetPublishers(ctx context.Context) ([]Publisher, error)
	GetPublisher(ctx context.Context, id int64) (Publisher, error)
	CreatePublisher(ctx context.Context, publisher *Publisher) error
	UpdatePublisher(ctx context.Context, id int64, publisher *Publisher) error
	DeletePublisher(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	// seller
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	CreateCategory(ctx context.Context, category *Category) error
	UpdateCategory(ctx context.Context, id int64, category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type BookRepository interface {
	// seller
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, id int64, book *Book) error
	DeleteBook(ctx context.Context, id int64) error
}

// Usecase
type PublisherUsecase interface {
	GetPublishers(ctx context.Context) ([]Publisher, error)
	GetPublisher(ctx context.Context, id int64) (Publisher, error)
	CreatePublisher(ctx context.Context, publisher *Publisher) error
	UpdatePublisher(ctx context.Context, id int64, publisher *Publisher) error
	DeletePublisher(ctx context.Context, id int64) error
}

type CategoryUsecase interface {
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	CreateCategory(ctx context.Context, category *Category) error
	UpdateCategory(ctx context.Context, id int64, category *Category) error
	DeleteCategory(ctx context.Context, id int64) error
}

type BookUsecase interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, id int64, book *Book) error
	DeleteBook(ctx context.Context, id int64) error
}
