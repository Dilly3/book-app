package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/dilly3/book-app/models"
	utils "github.com/dilly3/book-app/utils"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Validate *validator.Validate
	Client   *mongo.Client
	RWMutex  *sync.RWMutex
}

func (m Mongo) col(collectionName string) *mongo.Collection {
	return m.Client.Database("bookDB").Collection(collectionName)
}

func MongoDBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDB := os.Getenv("MONGODB_URL")
	if MongoDB == "" || len(MongoDB) < 1 {
		MongoDB = "mongodb://localhost:27017/bookDb"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
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
	book.ID = utils.GenerateRandomID()

	_, insertErr := m.col(models.BOOK_COLLECTION).InsertOne(ctx, book)
	if insertErr != nil {
		return nil, insertErr
	}
	return book, nil
}

func (m Mongo) GetBook(id string) (book *models.Book, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filterQuery := bson.M{
		"_id": id,
	}
	errdB := m.col(models.BOOK_COLLECTION).FindOne(ctx, filterQuery).Decode(&book)
	if errdB != nil {
		return nil, errdB
	}
	return book, nil
}

func (m Mongo) IsBookInStore(name string, author string) (bool, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{
		"title":  name,
		"author": author,
	}
	count, err := m.col(models.BOOK_COLLECTION).CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m Mongo) UpdateBook(id string, book *models.Book) (bk *models.Book, err error) {
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
	m.RWMutex.Lock()
	_ = m.col(models.BOOK_COLLECTION).FindOneAndUpdate(ctx, filter, updateQuery)
	m.RWMutex.Unlock()
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

func (m Mongo) DeleteBook(id string) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{
		"_id": id,
	}
	_, err = m.col(models.BOOK_COLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
