package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"lain.sceptix.net"
	"lain.sceptix.net/web"
)

func main() {

	dsn := "postgresql://nithin:REVEAL_PASSWORD@sunset-worm-8099.8nk.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	//The flag creation for the address
	var (
		addr       string
		sqlAddr    string
		sessionKey string
	)

	fs := flag.NewFlagSet("lain", flag.ExitOnError)
	fs.StringVar(&addr, "addr", ":4000", "HTTP service address")
	fs.StringVar(&sqlAddr, "sql-addr", "postgresql://root@PenPen:26257/defaultdb?sslmode=disable", "SQL address")
	fs.StringVar(&sessionKey, "session-key", "2p5TgHr7qJzWs4vY", "Session Key")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags %w", err)
	}

	//making a database connection
	db, err := sql.Open("postgres", sqlAddr)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping db :%w", err)
	}

	if err := lain.MigrateSQL(context.Background(), db); err != nil {
		return fmt.Errorf("migrate sql: %w", err)
	}

	queries := lain.New(db)
	svc := &lain.Service{
		Queries: queries,
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
	handler := &web.Handler{
		Logger:     logger,
		Service:    svc,
		SessionKey: []byte(sessionKey),
	}

	//now creating a server and passing the Handler
	srv := &http.Server{
		Handler: handler,
		//The address we take that from the flag
		Addr: addr,
	}

	defer srv.Close()

	//printing a log message here
	logger.Printf("listening on %s", addr)

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve error : %w", err)
	}

	return nil

}
