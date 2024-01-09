// main.go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"lain.sceptix.net"
	"lain.sceptix.net/web"
)

var db *sql.DB

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	setupDB()

	// Vercel provides the PORT environment variable for the HTTP service address
	addr := ":" + os.Getenv("PORT")

	queries := lain.New(db)
	svc := &lain.Service{
		Queries: queries,
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
	handler := &web.Handler{
		Logger:     logger,
		Service:    svc,
		SessionKey: []byte(os.Getenv("SESSION_KEY")), // Use environment variable for session key

	}

	// Creating an HTTP server
	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	defer srv.Close()

	// Printing a log message here
	logger.Printf("listening on %s", addr)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve error: %w", err)
	}

	return nil
}

func setupDB() {
	// Making a database connection
	var err error
	db, err = sql.Open("postgres", os.Getenv("SQL_ADDR"))
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database:", err)
	}

	if err := lain.MigrateSQL(context.Background(), db); err != nil {
		log.Fatal("failed to migrate SQL:", err)
	}
}
