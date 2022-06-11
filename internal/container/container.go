package container

import (
	"context"

	"github.com/roger.russel/90poe/internal/flags"
	"github.com/roger.russel/90poe/pkg/component/db"
	"github.com/roger.russel/90poe/pkg/component/streamer"
	"github.com/roger.russel/90poe/pkg/domain/port"
)

type Components struct {
	DB       db.Inter
	Streamer streamer.Inter
}

type Services struct {
	Port port.Inter
}

type Dependency struct {
	Cmp Components
	Srv Services
}

func New(ctx context.Context, flags flags.Flags) (context.Context, *Dependency, error) {
	jsonConfig := streamer.JSONConfig{
		BufferParserSize: flags.ParserBufferSize,
	}

	cmp := &Components{
		DB:       db.New(ctx),
		Streamer: streamer.NewJSON(jsonConfig),
	}

	repoConf := port.RepositoryConfig{
		StreamFile: flags.File,
	}

	portRepo := port.NewRepository(ctx, repoConf, cmp.Streamer, cmp.DB)
	port := port.New(ctx, portRepo)

	dep := Dependency{
		Cmp: Components{},
		Srv: Services{
			Port: port,
		},
	}

	return ctx, &dep, nil
}
