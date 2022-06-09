package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cancel context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go handleSystemCall(ctx, sig, cancel)

	log.Println("server is starting")

	// ToDO Add Run Server Here

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
