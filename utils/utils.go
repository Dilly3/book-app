package util

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if port == "" {
		port = "8080"
	}
	return port
}

func StartServer(r *gin.Engine) *http.Server {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := &http.Server{
		Addr:    "127.0.0.1" + GetPortFromEnv(),
		Handler: r,
	}
	return server
}
