package container

import (
	"context"

	"github.com/roger.russel/90poe/pkg/domain/port"
)

type Components struct {
}

type Services struct {
	Port port.Inter
}

type Dependency struct {
	Cmp Components
	Srv Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	port := port.New(ctx, nil)

	dep := Dependency{
		Cmp: Components{},
		Srv: Services{
			Port: port,
		},
	}

	return ctx, &dep, nil
}
