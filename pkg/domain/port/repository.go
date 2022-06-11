package port

import (
	"context"
	"fmt"

	"github.com/roger.russel/90poe/pkg/component/db"
	"github.com/roger.russel/90poe/pkg/component/streamer"
)

var _ RepositoryInter = (*Repository)(nil)

type RepositoryInter interface {
	Upsert(ctx context.Context, key string, data []byte)
	Stream(ctx context.Context, ch chan streamer.Data) error
}

type Repository struct {
	conf RepositoryConfig
	db   db.Inter
	stre streamer.Inter
}

type RepositoryConfig struct {
	StreamFile string
}

func NewRepository(ctx context.Context, conf RepositoryConfig, stre streamer.Inter, d db.Inter) *Repository {
	return &Repository{
		conf: conf,
		stre: stre,
		db:   d,
	}
}

func (r *Repository) Stream(ctx context.Context, ch chan streamer.Data) error {
	if err := r.stre.StreamFile(ctx, r.conf.StreamFile, ch); err != nil {
		return fmt.Errorf("error on stream: %w", err)
	}

	return nil
}

func (r *Repository) Upsert(ctx context.Context, key string, data []byte) {
	r.db.Upsert(ctx, key, data)
}

func (r *Repository) Table(ctx context.Context) db.KeyDB {
	return r.db.Table(ctx)
}
