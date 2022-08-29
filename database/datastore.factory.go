package database

import "github.com/go-playground/validator"

var mongoBK = MongoDBBKinstance()
var mongoUSR = MongoDBUSRinstance()
var Validate = validator.New()

func NewMongoBK() DataStore {
	return &MongoBK{
		Validate: Validate,
		Client:   mongoBK,
	}
}

func NewMongoUSR() UserStore {
	return &MongoUSR{
		Validate: Validate,
		Client:   mongoUSR,
	}
}
