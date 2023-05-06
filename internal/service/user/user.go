package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/utils"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Register(ctx context.Context, user Create) (string, error) {
	id := uuid.NewString()
	password, err := utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return "", err
	}
	create := Create{
		ID:        id,
		FullName:  user.FullName,
		Password:  string(password),
		Avatar:    user.Avatar,
		Role:      user.Role,
		BirthDate: user.BirthDate,
		Phone:     user.Phone,
		Position:  user.Position,
	}

	if err := s.repo.Create(ctx, create); err != nil {
		return "", err
	}

	return id, err
}
