package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ctrox/zeropod/manager"
)

var (
	metricsAddr = flag.String("metrics-addr", ":8080", "address of the metrics server")
)

func main() {
	flag.Parse()
	slog.Info("starting manager", "metrics-addr", *metricsAddr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := manager.AttachRedirectors(ctx); err != nil {
		slog.Error("attaching redirectors", "err", err)
		os.Exit(1)
	}

	server := &http.Server{Addr: *metricsAddr}
	http.HandleFunc("/metrics", manager.Handler)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Error("serving metrics", "err", err)
				os.Exit(1)
			}
		}
	}()

	<-ctx.Done()
	slog.Info("stopping manager")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("shutting down server", "err", err)
	}
}
