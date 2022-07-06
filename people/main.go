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

func main() {
	l := hclog.Default()
	err := run(&l)
	if err != nil {
		l.Error("Error shutting down the server", err)
		os.Exit(1)
	}
}

func run(l *hclog.Logger) error {
	logger := *l
	r := mux.NewRouter()

	srv := http.Server{
		Addr:         bindAddr,
		Handler:      r,
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("Server interrupted:", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Info("Gracefully shutting down the server:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}
