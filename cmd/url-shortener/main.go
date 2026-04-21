package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	logger.Info("starting url-shortener", slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	logger.Info("db started")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.EnvLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		//fmt.Fprintf(os.Stderr, "error: unknown environment: %s\n", env)
		fmt.Fprintf(os.Stderr, "%s unknown environment: %s\n", time.Now(), env)
		os.Exit(1)
	}

	return logger
}
