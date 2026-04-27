package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/mhdarifsetiawan/semantic-pantau-gizi/config"
	"github.com/mhdarifsetiawan/semantic-pantau-gizi/internal/handler"
	"github.com/mhdarifsetiawan/semantic-pantau-gizi/internal/repository"
	"github.com/mhdarifsetiawan/semantic-pantau-gizi/internal/service"
	"github.com/mhdarifsetiawan/semantic-pantau-gizi/pkg/logger"
)

func main() {
	cfg := config.Load()

	log := logger.New(logger.Config{
		Level:       cfg.Log.Level,
		Pretty:      cfg.Log.Pretty,
		ServiceName: "semantic-pantau-gizi",
	})

	userRepo := repository.NewUserMemoryRepository()
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	r := chi.NewMux()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.Recoverer)

	api := humachi.New(r, huma.DefaultConfig("Semantic Pantau Gizi API", "1.0.0"))

	handler.RegisterUserRoutes(api, userHandler)

	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("server started", map[string]any{
			"port": cfg.App.Port,
			"env":  cfg.App.Env,
			"docs": "http://localhost:" + cfg.App.Port + "/docs",
		})
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server failed to start", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown", err)
	}

	log.Info("server stopped")
}
