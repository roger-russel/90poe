package port

import (
	"context"
	"fmt"
	"log"

	"github.com/roger.russel/90poe/pkg/component/streamer"
)

var _ Inter = (*Service)(nil)

type Inter interface {
	Run(ctx context.Context) error
}

type Service struct {
	repo   RepositoryInter
	chData chan streamer.Data
}

func New(ctx context.Context, repo RepositoryInter) *Service {
	return &Service{
		repo:   repo,
		chData: make(chan streamer.Data),
	}
}

func (s *Service) Run(ctx context.Context) error {
	go s.upsertListener(ctx)

	if err := s.repo.Stream(ctx, s.chData); err != nil {
		return fmt.Errorf("error streaming: %w", err)
	}

	return nil
}

func (s *Service) upsertListener(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case d := <-s.chData:
			log.Println("received key:", string(d.KeyName))
			s.repo.Upsert(ctx, string(d.KeyName), d.KeyContent)
		}
	}
}
