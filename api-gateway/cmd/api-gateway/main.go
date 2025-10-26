package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/pujidjayanto/choochoohub/api-gateway/bootstrap"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server, err := bootstrap.NewApplicationServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	go func() {
		log.Printf("server running on %s", server.Port)
		if err := server.App.Listen(server.Port); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.App.ShutdownWithContext(shutdownCtx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Info("server stopped, resource cleaned up")
}
