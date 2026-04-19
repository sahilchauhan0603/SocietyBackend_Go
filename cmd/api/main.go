package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	config "github.com/sahilchauhan0603/society/internal/config"

	database "github.com/sahilchauhan0603/society/internal/database"
	"github.com/sahilchauhan0603/society/internal/http/middleware"
	routes "github.com/sahilchauhan0603/society/internal/http/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := database.DatabaseConnector(cfg); err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	router := mux.NewRouter()
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.SecurityHeadersMiddleware)
	router.Use(middleware.CORSMiddleware(cfg.Server.AllowedOrigins))
	router.Use(middleware.TraceResponseMiddleware)

	routes.InitializeRoutes(router)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("server is running on port %s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-shutdownSignal
	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("server stopped gracefully")
}
