package config

import (
	"fmt"
	"log"
	"net/http"
	"winartodev/book-store-be/delivery"
	"winartodev/book-store-be/handler"
	"winartodev/book-store-be/logger"
	"winartodev/book-store-be/repository"
	"winartodev/book-store-be/usecase"

	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

func NewConfig() Config {
	var cfg Config
	gotenv.Load(".env")
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Serve() {
	cfg := NewConfig()

	// db, err := NewMysql(&cfg)
	db, err := NewPostgres(&cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger.Init()

	categoryRepo := repository.NewMysqlCategory(db)
	categoryUsecase := usecase.NewCategoryUsecase(&usecase.CategoryRepository{CategoryRepo: categoryRepo})
	categoryHander := delivery.NewCategoryHandler(categoryUsecase, cfg.BookStoreUsername, cfg.BookStorePassword)

	publisherRepo := repository.NewMysqlPublisher(db)
	publisherUsecase := usecase.NewPublihserUsecase(&usecase.PublisherRepository{PublisherRepo: publisherRepo})
	publisherHandler := delivery.NewPublisherHandler(publisherUsecase, cfg.BookStoreUsername, cfg.BookStorePassword)

	bookRepo := repository.NewMysqlBook(db)
	bookUsecase := usecase.NewBookUsecase(&usecase.BookRepository{BookRepo: bookRepo})
	bookHandler := delivery.NewBookHandler(bookUsecase, cfg.BookStoreUsername, cfg.BookStorePassword)

	h := handler.NewHandler(&categoryHander, &publisherHandler, &bookHandler)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: h,
	}

	logger.Info(fmt.Sprintf("bookstore API run on %s", s.Addr), logger.Fields{})
	if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
		log.Fatal(serr)
	}
}
