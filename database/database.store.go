package database

import (
	"github.com/dilly3/book-app/models"
)

type Datastore interface {
	CreateBook(book *models.Book) (*models.Book, error)
}
