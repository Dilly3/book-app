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
	Status      string    `json:"status" bson:"status"`
	RentedBy    *UserInfo `json:"rented_by"  bson:"rented-by"`
	TimeRented  time.Time `json:"time_rented"  bson:"time-rented"`
	Created_at  time.Time `json:"created_at"  bson:"created-at"`
	Updated_at  time.Time `json:"updated_at" bson:"updated-at"`
}
