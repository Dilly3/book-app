package main

import (
	"log"
	"time"

	"github.com/dilly3/book-rental/database"
	"github.com/dilly3/book-rental/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	time.Sleep(time.Millisecond * 1500)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ginHandler := routes.MountGinHandler(routes.NewHandle(database.NewMongoUSR, database.NewMongoBK))
	server := routes.StartServer(ginHandler)
	done := make(chan error, 1)

	go routes.GracefulShutdown(done, server)

	log.Fatal(server.ListenAndServe())
	<-done

}
