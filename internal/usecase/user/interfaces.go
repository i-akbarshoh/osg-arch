package user

import (
	"context"
	"github.com/i-akbarshoh/osg-arch/internal/service/user"
)

type User interface {
	Register(ctx context.Context, user user.Create) (string, error)
}
