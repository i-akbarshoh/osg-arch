package user

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/attendance"
)

func (u *UseCase) CreateAttendance(ctx context.Context,at entity.UserAttendance) error {
	return u.a.CreateAttendance(ctx, attendance.Create{
		Type: at.Type,
		UserID: at.UserID,
		CreatedAt: at.CreatedAt,
	})
}

func (u *UseCase) ListAttendance(ctx context.Context, user_id string, ty string) (attendance.List, error){
	return u.a.ListAttendance(ctx, user_id, ty)
}