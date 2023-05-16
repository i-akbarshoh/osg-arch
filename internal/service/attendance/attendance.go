package attendance

import "context"

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateAttendance(ctx context.Context, c Create) error {
	return s.repo.Create(ctx, c)
}

func (s *service) ListAttendance(ctx context.Context, user_id string, t string) (List, error) {
	return s.repo.List(ctx, user_id, t)
}