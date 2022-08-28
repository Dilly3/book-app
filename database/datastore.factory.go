package database

import "github.com/go-playground/validator"

var Mongo = MongoDBinstance()
var Validate = validator.New()

func NewMongoBK() DataStore {
	return &MongoBK{
		Validate: Validate,
		Client:   Mongo,
	}
}

func NewMongoUSR() UserStore {
	return &MongoUSR{
		Validate: Validate,
		Client:   Mongo,
	}
}
