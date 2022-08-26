package main

import (
	"github.com/dilly3/book-app/database"
	"github.com/dilly3/book-app/routes"
	"log"
)

func main() {

	ginHandler := routes.MountGinHandler(routes.NewHandle(database.NewMongoDb))
	server := routes.StartServer(ginHandler)
	done := make(chan error, 1)

	go routes.GracefulShutdown(done, server)

	log.Fatal(server.ListenAndServe())
	<-done

}
