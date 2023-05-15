package task

import (
	"context"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateTask(ctx context.Context, c Create) (int, error) {
	c.ID = uuid.NewString()

	id, err := s.repo.Create(ctx, c)
	if err != nil {
		return 0, err
	}

	return id, nil
}