package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"moonshine-server/gen/moonshine/v1/moonshinev1connect"
	"moonshine-server/internal/config"
	"moonshine-server/internal/db"
	"moonshine-server/internal/service"
)

func main() {
	cfg := config.Load()

	// Connect to PostgreSQL
	ctx := context.Background()
	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to PostgreSQL")

	// Create service handlers
	authService := service.NewAuthService(pool)
	chatService := service.NewChatService(pool)
	channelService := service.NewChannelService(pool)
	userService := service.NewUserService(pool)

	// Set up HTTP mux with ConnectRPC handlers
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok"}`)
	})

	// Mount ConnectRPC service handlers
	path, handler := moonshinev1connect.NewAuthServiceHandler(authService)
	mux.Handle(path, handler)

	path, handler = moonshinev1connect.NewChatServiceHandler(chatService)
	mux.Handle(path, handler)

	path, handler = moonshinev1connect.NewChannelServiceHandler(channelService)
	mux.Handle(path, handler)

	path, handler = moonshinev1connect.NewUserServiceHandler(userService)
	mux.Handle(path, handler)

	// Configure server
	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		}
	}()

	log.Printf("Moonshine server listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
	log.Println("Server stopped")
}
