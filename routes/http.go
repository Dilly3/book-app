package routes

import (
	"time"

	controllers "github.com/dilly3/book-app/controller"
	"github.com/dilly3/book-app/database"

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
	mongodb.Client = controllers.DBinstance()
	handler := new(Handler)
	handler.store = mongodb
	router.POST("/books/create", handler.CreateBook())
	router.GET("/books/:book_id", handler.GetBook())
	router.PATCH("/editbook/:book_id", handler.UpdateBook())
	router.GET("/getallbooks", handler.GetAllBooks())

	return router
}

//
// incomingRoutes.DELETE("books/:book_id", controllers.DeleteBook())
// incomingRoutes.GET("books", controllers.GetAllBooks())
