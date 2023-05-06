package user

import (
	"context"
	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/user"
)

type UseCase struct {
	u User
}

func NewUseCase(u User) *UseCase {
	return &UseCase{
		u: u,
	}
}

func (u *UseCase) Register(ctx context.Context, ur entity.User) (string, error) {
	return u.u.Register(ctx, user.Create{
		FullName:  ur.FullName,
		Password:  ur.Password,
		Avatar:    ur.Avatar,
		Role:      ur.Role,
		BirthDate: ur.BirthDate,
		Phone:     ur.Phone,
		Position:  ur.Position,
	})
}
