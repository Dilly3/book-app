package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dilly3/book-app/routes"
)

func main() {
	ginHandler := routes.MountGinHandler()
	server := routes.StartServer(ginHandler)
	done := make(chan error, 1)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		fmt.Println("\nshutting down server")
		time.Sleep(time.Second * 5)
		ctx := context.Background()
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		done <- server.Shutdown(ctx)

	}()
	log.Fatal(server.ListenAndServe())
	log.Println(<-done)

}
