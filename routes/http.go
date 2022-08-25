package routes

import (
	"github.com/dilly3/book-app/database"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func MountServer() *gin.Engine {

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	mongodb := database.Mongo{}
	mongodb.Validate = validator.New()
	mongodb.Client = database.DBinstance()
	handler := new(Handler)
	handler.store = mongodb
	router.POST("/books/createbook", handler.CreateBook())
	router.GET("/books/getbook/:book_id", handler.GetBook())
	router.PATCH("/books/editbook/:book_id", handler.UpdateBook())
	router.GET("/books/getallbooks", handler.GetAllBooks())
	router.DELETE("/books/deletebook/:_id", handler.DeleteBook())

	return router
}
