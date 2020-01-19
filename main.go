package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bonggar/gorestapi/config"
	"github.com/bonggar/gorestapi/database"
	"github.com/bonggar/gorestapi/router"
)

func main() {
	// Load config
	config.Load()

	// Open connection to database
	database.SQLiteDBConnect()
	defer database.GetDB().Close()

	// Register routes
	routeHandler := router.Make()
	srv := &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: routeHandler,
	}

	// Run HTTP service
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Ready to serve
	log.Println("listen on", srv.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
