package database

import "github.com/go-playground/validator"

func NewMongoDb() DataStore {
	return &Mongo{
		Validate: validator.New(),
		Client:   MongoDBinstance(),
	}
}
