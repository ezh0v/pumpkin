package main

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezh0v/pumpkin/internal/app"
	"github.com/ezh0v/pumpkin/internal/pkg/postgres"
	"github.com/ezh0v/pumpkin/internal/server"
)

func main() {
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

	app := app.NewContext(os.Getenv("APP_VERSION"), database)

	server, err := server.New(app)
	if err != nil {
		slog.Error("server initialization failed", "error", err)
		return
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("server terminated with an error", "error", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(); err != nil {
		slog.Error("server shutdown failed", "error", err)
	}
}

func close(resources ...io.Closer) {
	for _, resource := range resources {
		if err := resource.Close(); err != nil {
			slog.Error("failed to close resource", "error", err)
		}
	}
}
