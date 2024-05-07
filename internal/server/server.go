package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prafitradimas/go-http-impl/internal/routes"
)

type ServerConfig struct {
	Host   string
	Port   int
	Logger *slog.Logger
}

func Run(ctx context.Context) error {
	logger := slog.Default()
	config := &ServerConfig{
		Host:   "localhost",
		Port:   8080,
		Logger: logger,
	}

	httpHandlers := routes.SetupHandlers(
		logger,
		nil,
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, strconv.Itoa(config.Port)),
		Handler: httpHandlers,
	}

	go func() {
		logger.Info(fmt.Sprintf("Listening on %s", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("Error listening and serving: %s \n", err))
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		// shutdown context
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		fmt.Println("Shutting down server...")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Error(fmt.Sprintf("Error shutting down http server: %s \n", httpServer.Addr))
		}
	}()
	wg.Wait()

	return nil
}
