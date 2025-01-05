package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func DBINIT() {
	err := InitDb()
	if err != nil {
		log.Fatal("Can not intiate the DB instace. Error: ", err)
	}
}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting the Application...")
	server := StartServer()

	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)
	DBINIT()
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}
