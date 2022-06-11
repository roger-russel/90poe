package container

import (
	"context"

	"github.com/roger.russel/90poe/internal/flags"
	"github.com/roger.russel/90poe/pkg/component/orm"
	"github.com/roger.russel/90poe/pkg/component/streamer"
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

func New(ctx context.Context, flags flags.Flags) (context.Context, *Dependency, error) {
	jsonConfig := streamer.JsonConfig{
		BufferParserSize: flags.ParserBufferSize,
	}

	portRepo := port.NewRepository(ctx, streamer.NewJson(jsonConfig), orm.New(ctx))
	port := port.New(ctx, portRepo)

	dep := Dependency{
		Cmp: Components{},
		Srv: Services{
			Port: port,
		},
	}

	return ctx, &dep, nil
}
