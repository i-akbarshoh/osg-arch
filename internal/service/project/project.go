package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
)

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context, c Create) (int, error) {
	return s.repo.Create(ctx, c)
}

func (s *service) List(ctx context.Context) (entity.List, error) {
	return s.repo.List(ctx)
}

func (s *service) Update(ctx context.Context, u Update) error {
	return s.repo.UpdateStatus(ctx, u)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}