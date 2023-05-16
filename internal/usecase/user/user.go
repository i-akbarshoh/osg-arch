package user

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/user"
)

type UseCase struct {
	a Attendance
	u User
}

func NewUseCase(u User, a Attendance) *UseCase {
	return &UseCase{
		u: u,
		a: a,
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

func (u *UseCase) Get(ctx context.Context, id string) (entity.User, error) {
	res := entity.User{}
	get, err := u.u.Get(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	res.ID = get.ID
	res.Avatar = get.Avatar
	res.BirthDate = get.BirthDate
	res.FullName = get.FullName
	res.Phone = get.Phone
	res.Position = get.Position
	res.Role = get.Role
	return res, err
}

func (u *UseCase) Update(ctx context.Context,us entity.User) error {
	return u.u.Update(ctx, user.Update{
		ID: us.ID,
		FullName: us.FullName,
		Avatar: us.Avatar,
		Role: us.Role,
		Phone: us.Phone,
		Position: us.Position,
	})
}

func (u *UseCase) Delete(ctx context.Context, ud entity.User) error {
	return u.u.Delete(ctx, user.Delete{
		ID: ud.ID,
		Password: ud.Password,
	})
}
