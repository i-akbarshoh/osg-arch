package user

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/service/attendance"
	"github.com/i-akbarshoh/osg-arch/internal/service/user"
)

type User interface {
	Register(ctx context.Context, user user.Create) (string, error)
	Login(context.Context, user.Get) error
	List(context.Context) (user.List, error)
	Get(context.Context, string) (user.Get, error)
	Update(context.Context, user.Update) error
	Delete(context.Context, user.Delete) error
}

type Attendance interface {
	CreateAttendance(context.Context, attendance.Create) error
	ListAttendance(context.Context, string, string) (attendance.List, error)
}