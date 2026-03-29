package main

import (
	"context"
	"log"
	"net/http"
	"time"

	httpHandlers "github.com/nathantotten/random-fullstack/backend/internal/http"
	"github.com/nathantotten/random-fullstack/backend/internal/platform"
	"github.com/nathantotten/random-fullstack/backend/internal/recipe"
)

func main() {
	// 1️⃣ Load config
	cfg, err := platform.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2️⃣ Initialize logger
	logger := platform.NewLogger(cfg)

	logger.Info("Starting Coffee Brew Journal API...")

	// 3️⃣ Connect to DB
	db, err := platform.NewDB(cfg)
	if err != nil {
		logger.Fatal("failed to connect to DB", "err", err)
	}
	defer db.Close()

	// 4️⃣ Initialize sqlc queries
	queries := platform.NewQueries(db)

	// 5️⃣ Initialize repositories
	recipeRepo := recipe.NewRepository(queries)

	// 6️⃣ Initialize services
	recipeService := recipe.NewService(recipeRepo)

	// 7️⃣ Initialize HTTP handlers
	recipeHandler := httpHandlers.NewRecipeHandler(recipeService)

	// 8️⃣ Setup router
	mux := http.NewServeMux()
	httpHandlers.RegisterRoutes(mux, recipeHandler) // central place for all routes

	// 9️⃣ Wrap with logging middleware
	handler := httpHandlers.LoggingMiddleware(mux, logger)

	// 10️⃣ Start server
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logger.Info("Server listening on " + cfg.ServerAddress)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server failed", "err", err)
	}

	// 11️⃣ Graceful shutdown (optional for sprint 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
