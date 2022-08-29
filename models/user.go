package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        string    `json:"_id"  bson:"_id"`
	UserName  string    `json:"username" bson:"username"`
	Email     string    `json:"email" bson:"email"`
	Password  *string   `json:"password" bson:"password"`
	Age       *int      `json:"age" bson:"age"`
	Role      string    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
}

type UserInfo struct {
	ID       string `json:"_id"  bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type JWTClaims struct {
	*jwt.RegisteredClaims
	ID    string `json:"id"`
	Email string `json:"email"`
}
type DtoClaims struct {
	PersonId string `json:"person_id"`
	Email    string `json:"email"`
}
