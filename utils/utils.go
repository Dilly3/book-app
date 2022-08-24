package util

import "go.mongodb.org/mongo-driver/bson/primitive"

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
