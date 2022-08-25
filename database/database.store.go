package database

import (
	"github.com/dilly3/book-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Datastore interface {
	AddBook(book *models.Book) (*models.Book, error)
	GetBook(id primitive.ObjectID) (book *models.Book, err error)
	UpdateBook(id primitive.ObjectID, book *models.Book) (bk *models.Book, err error)
	GetAllBooks() (books []*models.Book, err error)
	DeleteBook(id primitive.ObjectID) (err error)
}
