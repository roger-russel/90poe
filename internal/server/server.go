package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/roger.russel/90poe/internal/container"
)

func Run(ctx context.Context, cancel context.CancelFunc, debs *container.Dependency) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go handleSystemCall(ctx, sig, cancel)

	log.Println("server is starting")

	if err := run(ctx, &debs.Srv, cancel); err != nil {
		log.Println("critical error on start", err)
		cancel()
	}

	<-ctx.Done()
	log.Println("server is shutting down")
}

func handleSystemCall(ctx context.Context, sig chan os.Signal, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		return
	case <-sig:
		log.Println("system call server to shutdown")
		cancel()

		return
	}
}

func run(ctx context.Context, srv *container.Services, cancel context.CancelFunc) error {
	if err := srv.Port.Run(ctx); err != nil {
		return errors.Wrap(err, "error running port server")
	}
	cancel() // start shutdown

	return nil
}
