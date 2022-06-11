package main

import (
	"context"
	"fmt"
	"log"

	"github.com/roger.russel/90poe/internal/container"
	"github.com/roger.russel/90poe/internal/flags"
	"github.com/roger.russel/90poe/internal/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		if r := recover(); r != nil {
			log.Println("panic recovered on main:", r)
		}
	}()

	ctx, dep, err := container.New(ctx, flags.Load())
	if err != nil {
		panic(err) // panic is acceptable only on main
	}
	if err = run(ctx, cancel, dep); err != nil {
		panic(err) // panic is acceptable only on main
	}
}

func run(ctx context.Context, cancel context.CancelFunc, dep *container.Dependency) error {
	if err := server.Run(ctx, cancel, dep); err != nil {
		return fmt.Errorf("error running server: %w", err)
	}
	return nil
}
