package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dilly3/library-manager/models"
	utils "github.com/dilly3/library-manager/utils"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUSR struct {
	Validate *validator.Validate
	Client   *mongo.Client
	RWMutex  *sync.RWMutex
}

func (m MongoUSR) colUSR() *mongo.Collection {
	return m.Client.Database("userDB").Collection(models.USER_COLLECTION)
}
func MongoDBUSRinstance() *mongo.Client {

	MongoDB := os.Getenv("MONGODBUSR_URL")
	if MongoDB == "" || len(MongoDB) < 1 {
		MongoDB = "mongodb://localhost:27017/userdb"
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

	fmt.Println("Connected to MongoDBUSR")
	return client
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
func (m MongoUSR) GetUserByEmail(email string) (*models.User, error) {
	var user = &models.User{}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.M{
		"email": email,
	}
	err := m.colUSR().FindOne(ctx, filter).Decode(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
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
