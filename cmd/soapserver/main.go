package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cbrapplication/soapserver"
)

func main() {
	port := "8080"
	server := soapserver.NewSOAPServer(port)
	if server == nil {
		panic("Cant load")
	}

	go func() {

		log.Println("Starting Server")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	waitForShutdown(server)

}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
