package routes

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	utils "github.com/dilly3/library-manager/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
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

	router.POST("/admin/addbook", handler.CreateBook())
	router.GET("/admin/getbook/:book_id", handler.GetBook())
	router.PATCH("/admin/editbook/:book_id", handler.UpdateBook())
	router.GET("/user/getallbooks", handler.GetAllBooks())
	router.GET("/", handler.Home())
	router.DELETE("/admin/deletebook/:_id", handler.DeleteBook())
	router.POST("/user/signup", handler.UserSignUp())
	router.POST("/user/signin", handler.UserLogin())

	return router
}

func StartServer(r *gin.Engine) *http.Server {
	port := utils.GetPortFromEnv()
	if port == "" {
		port = ":8080"
	}

	server := &http.Server{
		Addr:    "127.0.0.1" + port,
		Handler: r,
	}
	return server
}

func GracefulShutdown(done chan error, server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("\nshutting down server")
	time.Sleep(time.Millisecond * 500)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	fmt.Print("B")
	time.Sleep(time.Millisecond * 500)
	fmt.Print("Y")
	time.Sleep(time.Millisecond * 500)
	fmt.Print("E")
	time.Sleep(time.Millisecond * 500)
	fmt.Print("!")
	time.Sleep(time.Millisecond * 500)
	fmt.Print(" ")
	time.Sleep(time.Millisecond * 500)
	fmt.Print("\n")
	done <- server.Shutdown(ctx)

}
