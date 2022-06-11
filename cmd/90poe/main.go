package main

import (
	"context"
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

	server.Run(ctx, cancel, dep)
}
