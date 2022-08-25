package main

import (
	"log"

	"github.com/dilly3/book-app/database"
	"github.com/dilly3/book-app/routes"
)

func main() {
	ginHandler := routes.MountGinHandler(routes.NewHandle(database.NewMongoDb))
	server := routes.StartServer(ginHandler)
	done := make(chan error, 1)

	go routes.GracefulShutdown(done, server)

	log.Fatal(server.ListenAndServe())
	<-done

}
