package task

import (
	"context"
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
	id, err := s.repo.Create(ctx, c)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) Update(ctx context.Context, u Update) error {
	return s.repo.UpdateStatus(ctx, u)
}

func (s *service) Get(ctx context.Context, id int) (Get, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) ListTasks(ctx context.Context) (List, error) {
	return s.repo.List(ctx)
}

func (s *service) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}