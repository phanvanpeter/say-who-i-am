package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddr = ":9090"

var logger hclog.Logger

func main() {
	logger = hclog.Default()
	err := run()
	if err != nil {
		logger.Error("Error shutting down the server", err)
		os.Exit(1)
	}
}

// run starts the server for people API
func run() error {
	r := mux.NewRouter()

	srv := newServer(r)
	startServer(srv)

	waitForSignal(logger)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

// newServer initializes and returns a new server
func newServer(r *mux.Router) *http.Server {
	return &http.Server{
		Addr:         bindAddr,
		Handler:      r,
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}
}

// startServer listens on the TCP network address and then calls Serve to handle requests on incoming connections.
// A new goroutine is initialized for this process.
func startServer(srv *http.Server) {
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("Server interrupted:", err)
			os.Exit(1)
		}
	}()
}

// waitForSignal waits for the os.Interrupt or os.Kill signal. Until then, the main goroutine is blocked.
// This prevents from the immediate shutdown of the server.
func waitForSignal(logger hclog.Logger) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Info("Gracefully shutting down the server:", sig)
}