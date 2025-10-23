package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pujidjayanto/choochoohub/user-api/bootstrap"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server, cleanup, err := bootstrap.NewApplicationServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	go func() {
		log.Printf("server running on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := cleanup(shutdownCtx); err != nil {
		log.Fatalf("cleanup failed: %v", err)
	}

	log.Println("server stopped, resource cleaned up")
}
