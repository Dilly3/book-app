package database

import (
	"context"
	"fmt"

	"time"

	"github.com/dilly3/book-app/models"
	"github.com/go-playground/validator"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Validate *validator.Validate
	Client   *mongo.Client
}

func (m Mongo) col(collectionName string) *mongo.Collection {
	return m.Client.Database("bookDB").Collection(collectionName)
}
func (m Mongo) AddBook(book *models.Book) (*models.Book, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	validationErr := m.Validate.Struct(book)
	if validationErr != nil {
		return nil, validationErr
	}

	book.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.ID = primitive.NewObjectID()

	_, insertErr := m.col(models.BOOK_COLLECTION).InsertOne(ctx, book)
	if insertErr != nil {
		return nil, insertErr
	}
	return book, nil
}

func (m Mongo) GetBook(id primitive.ObjectID) (book *models.Book, err string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filterQuery := bson.M{
		"_id": id,
	}
	errdB := m.col(models.BOOK_COLLECTION).FindOne(ctx, filterQuery).Decode(&book)
	if errdB != nil {
		return nil, fmt.Sprintf("error while getting book : %s", errdB.Error())
	}
	return book, ""
}

func (m Mongo) UpdateBook(id primitive.ObjectID, book *models.Book) (bk *models.Book, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": id,
	}

	var updateObj = bson.M{}

	if book.Author != nil {
		updateObj["author"] = book.Author

	}
	if book.Title != nil {
		updateObj["title"] = book.Title
	}
	if book.Description != nil {
		updateObj["description"] = book.Description
	}

	updateObj["updated_at"] = time.Now().Format(time.RFC3339)
	updateObj["_id"] = id
	updateQuery := bson.M{
		"$set": updateObj,
	}
	var bookCollection *mongo.Collection = m.Client.Database("bookDB").Collection(models.BOOK_COLLECTION)
	_, err = bookCollection.UpdateOne(ctx, filter, updateQuery)
	if err != nil {
		return nil, err
	}
	return book, nil

}

func (m Mongo) GetAllBooks() (books []*models.Book, err error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := m.col(models.BOOK_COLLECTION).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	books = make([]*models.Book, cursor.RemainingBatchLength())
	if err := cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, nil
}
