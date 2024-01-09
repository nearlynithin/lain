// api/index.go
package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"lain.sceptix.net"
	"lain.sceptix.net/web"
)

var db *sql.DB

func init() {
	setupDB()
}

func setupDB() {
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

func Handler(w http.ResponseWriter, r *http.Request) {
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

	// Serve HTTP using the custom handler
	handler.ServeHTTP(w, r)
}
