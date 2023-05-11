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

func (u *UseCase) Login(ctx context.Context, ug entity.User) (err error) {
	err = u.u.Login(ctx, user.Get{
		ID: ug.ID,
		Password: ug.Password,
	})

	return
}

func (u *UseCase) List(ctx context.Context) (entity.UserList, error) {
	res := entity.UserList{}
	list, err := u.u.List(ctx)
	if err != nil {
		return entity.UserList{}, err
	}
	for _, v := range list.L{
		res.L = append(res.L, entity.User{
			ID: v.ID,
			FullName: v.FullName,
			Password: v.Password,
			Avatar: v.Avatar,
			Role: v.Role,
			BirthDate: v.BirthDate,
			Phone: v.Phone,
			Position: v.Position,
		})
	}
	res.Count = list.Count

	return res, nil
}
