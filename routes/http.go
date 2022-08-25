package routes

import (
	utils "github.com/dilly3/book-app/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func MountGinHandler() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler := NewHandle()

	router.POST("/books/createbook", handler.CreateBook())
	router.GET("/books/getbook/:book_id", handler.GetBook())
	router.PATCH("/books/editbook/:book_id", handler.UpdateBook())
	router.GET("/books/getallbooks", handler.GetAllBooks())
	router.DELETE("/books/deletebook/:_id", handler.DeleteBook())

	return router
}

func StartServer(r *gin.Engine) *http.Server {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := &http.Server{
		Addr:    "127.0.0.1" + utils.GetPortFromEnv(),
		Handler: r,
	}
	return server
}
