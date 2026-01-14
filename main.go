package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezh0v/pumpkin/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	server, err := server.New()
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
