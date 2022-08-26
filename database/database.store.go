package database

import (
	"github.com/dilly3/book-app/models"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type DataStore interface {
	AddBook(book *models.Book) (*models.Book, error)
	GetBook(id string) (book *models.Book, err error)
	UpdateBook(id string, book *models.Book) (bk *models.Book, err error)
	GetAllBooks() (books []*models.Book, err error)
	DeleteBook(id string) (err error)
	IsBookInStore(name string, author string) (bool, error)
}
