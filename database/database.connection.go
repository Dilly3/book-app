package database

import (
	"context"

	"time"

	"github.com/dilly3/book-app/models"
	"github.com/go-playground/validator"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Validate *validator.Validate
	Client   *mongo.Client
}

func (m *Mongo) CreateBook(book *models.Book, collectionName string) (*models.Book, error) {
	var bookCollection *mongo.Collection = m.Client.Database("book-DB").Collection(collectionName)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	validationErr := m.Validate.Struct(book)
	if validationErr != nil {
		return nil, validationErr
	}

	book.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.ID = primitive.NewObjectID()

	_, insertErr := bookCollection.InsertOne(ctx, book)
	if insertErr != nil {
		return nil, insertErr
	}
	return book, nil
}
