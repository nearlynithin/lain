// shared/setupdb.go
package shared

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"lain.sceptix.net"
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
