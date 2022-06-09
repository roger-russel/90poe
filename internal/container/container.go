package container

import (
	"context"
)

type Components struct {
}

type Services struct {
}

type Dependency struct {
	Cmp Components
	Srv Services
}

func New(ctx context.Context) (context.Context, *Dependency, error) {
	dep := Dependency{
		Cmp: Components{},
		Srv: Services{},
	}

	return ctx, &dep, nil
}
