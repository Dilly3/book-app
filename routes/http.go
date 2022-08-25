package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MountRouter() *gin.Engine {

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

	handler := NewHandler()

	router.POST("/books/createbook", handler.CreateBook())
	router.GET("/books/getbook/:book_id", handler.GetBook())
	router.PATCH("/books/editbook/:book_id", handler.UpdateBook())
	router.GET("/books/getallbooks", handler.GetAllBooks())
	router.DELETE("/books/deletebook/:_id", handler.DeleteBook())

	return router
}
