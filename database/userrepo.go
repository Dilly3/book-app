package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/dilly3/book-rental/models"
	utils "github.com/dilly3/book-rental/utils"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUSR struct {
	Validate *validator.Validate
	Client   *mongo.Client
	RWMutex  *sync.RWMutex
}

func (m MongoUSR) colUSR() *mongo.Collection {
	return m.Client.Database("bookDB").Collection(models.USER_COLLECTION)
}
func (m MongoUSR) CheckUserByEmail(email string) bool {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{
		"email": email,
	}
	count, err := m.colUSR().CountDocuments(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	if count > 0 {
		return true
	}
	return false
}

func (m MongoUSR) CreateUser(user *models.User) (*models.User, error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	validationErr := m.Validate.Struct(user)
	if validationErr != nil {
		return nil, validationErr
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = utils.GenerateRandomIDUSR()

	_, insertErr := m.colUSR().InsertOne(ctx, user)
	if insertErr != nil {
		return nil, insertErr
	}
	return user, nil
}
