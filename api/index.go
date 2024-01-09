// api/index.go
package api

import (
	"log"
	"net/http"
	"os"

	"shared" // Import shared functionality

	"lain.sceptix.net"
	"lain.sceptix.net/web"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	queries := lain.New(shared.DB) // Use the shared database instance
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
