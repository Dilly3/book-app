package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}
type SuccessResponse struct {
	Code    int         `json:"code"`
	Object  interface{} `json:"object"`
	Message interface{} `json:"message"`
}

func GetPrimitiveObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func GenerateObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}
func GetPortFromEnv() string {
	port := os.Getenv("PORT")
	if port == "" || len(port) < 1 {
		port = "8080"
	}
	return port
}
