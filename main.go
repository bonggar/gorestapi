package main

import (
	"log"
	"net/http"

	"github.com/bonggar/gorestapi/database"
	"github.com/bonggar/gorestapi/router"
)

func main() {
	// Open connection to database
	database.SQLiteDBConnect()
	defer database.GetDB().Close()

	// Register routes
	routeHandler := router.Make()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routeHandler,
	}

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	// Ready to serve
	log.Println("listen on", srv.Addr)
}
