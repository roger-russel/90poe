package port

import (
	"context"
)

type Inter interface {
	Run(ctx context.Context) error
}

type Service struct {
	repo RepositoryInter
}

func New(ctx context.Context, repo RepositoryInter) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Run(ctx context.Context) error {
	return nil
}
