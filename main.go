// main.go

package main

import (
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"github.com/Nageshbansal/web-crawler/pkg/api"
	"net/http"
)

func main() {
	// Initialize the HTTP router
	router := api.NewRouter()

	// Server address and port
	const serverAddr = ":8080"

	// Create a new server
	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Start the server
	logger.Infof("Starting server on port %s", serverAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", serverAddr, err)
	}

}
