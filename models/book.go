package models

import (
	"time"

	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          string    `bson:"_id"`
	Author      *string   `json:"author" validate:"required" bson:"author"`
	Title       *string   `json:"title" validate:"required" bson:"title" unique:"true"`
	Description *string   `json:"description" validate:"required"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
