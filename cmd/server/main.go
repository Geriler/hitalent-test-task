package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Geriler/hitalent/internal/app"
	"github.com/Geriler/hitalent/internal/config"
)

func main() {
	cfg := config.MustGet()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	a := app.NewApp(logger)

	go func() {
		if err := a.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("app start error", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.TimeoutStop)
	defer shutdownCancel()

	if err := a.Stop(shutdownCtx); err != nil {
		logger.Error("app stop error", "err", err)
	}
}
