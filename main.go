package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/postgres"
	"github.com/ezh0v/pumpkin/internal/server"
)

func main() {
	local, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		slog.Error("failed to load timezone", "error", err)
		return
	}
	time.Local = local

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	database, err := postgres.New(os.Getenv("DATABASE_CONNECT"))
	if err != nil {
		slog.Error("postgres initialization failed", "error", err)
		return
	}
	defer close(database)

	app := app.New(os.Getenv("APP_VERSION"), database)

	server, err := server.New(app)
	if err != nil {
		slog.Error("server initialization failed", "error", err)
		return
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("server terminated with an error", "error", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
}

func close(resource io.Closer) {
	if err := resource.Close(); err != nil {
		slog.Error("failed to close resource", "error", err)
	}
}
