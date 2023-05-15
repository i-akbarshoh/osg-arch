package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/i-akbarshoh/osg-arch/internal/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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

func (s *service) Login(ctx context.Context, get Get) error {
	res, err := s.repo.Get(ctx, get.ID)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(get.Password)); err != nil {
		return err
	}

	return nil
}

func (s *service) List(ctx context.Context) (List, error) {
	return s.repo.List(ctx)
}

func (s *service) Get(ctx context.Context, id string) (Get, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) Update(ctx context.Context, u Update) error {
	return s.repo.Update(ctx, u)
}

func (s *service) Delete(ctx context.Context, d Delete) error {
	if err := s.Login(ctx, Get{
		ID: d.ID,
		Password: d.Password,
	}); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, d); err != nil {
		return err
	}

	return nil
}