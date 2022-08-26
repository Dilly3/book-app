package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	utils "github.com/dilly3/book-app/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func MountGinHandler(handler *Handle) *gin.Engine {

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

func GracefulShutdown(done chan error, server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("\nshutting down server")
	time.Sleep(time.Second * 3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	fmt.Print("B")
	time.Sleep(time.Second * 1)
	fmt.Print("Y")
	time.Sleep(time.Second * 1)
	fmt.Print("E")
	time.Sleep(time.Second * 1)
	fmt.Print(" ")
	time.Sleep(time.Second * 1)
	fmt.Print("\n")
	done <- server.Shutdown(ctx)

}
