package comment

import "context"

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context,c Create) (id int, err error) {
	return s.repo.Create(ctx, c)
}

func (s *service) ListComments(ctx context.Context, task_id int) (List, error) {
	return s.repo.List(ctx, task_id)
}

func (s *service) DeleteComment(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}