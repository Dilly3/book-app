package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/dilly3/book-rental/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	bcrypto "golang.org/x/crypto/bcrypt"
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
func GenerateRandomID() string {
	b := make([]byte, 9)

	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s%x", "BK631-", b)
}
func GenerateRandomIDUSR() string {
	b := make([]byte, 9)

	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s%x", "USR-", b)
}
func RandomBigInt() int64 {
	bign, err1 := rand.Int(rand.Reader, big.NewInt(5000000))
	bigx, err := rand.Int(rand.Reader, big.NewInt(70000))
	if err != nil {
		log.Fatal(err)
	}
	if err1 != nil {
		log.Fatal(err)
	}
	return bign.Int64() + bigx.Int64()
}

func GetPresentTime() time.Time {
	time, err := time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
	if err != nil {
		log.Fatal(err)
	}
	return time
}

func EncryptPassword(password *string) *string {
	byte, err := bcrypto.GenerateFromPassword([]byte(*password), models.DEFAULT_COST)
	if err != nil {
		log.Fatal(err)
	}
	pass := string(byte)
	return &pass

}
