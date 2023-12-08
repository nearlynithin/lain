package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"lain.sceptix.net/web"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	//The flag creation for the address
	var addr string
	fs := flag.NewFlagSet("lain", flag.ExitOnError)
	fs.StringVar(&addr, "addr", ":8080", "HTTP service address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags %w", err)
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
	handler := &web.Handler{
		Logger: logger,
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

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve error : %w", err)
	}

	return nil

}
