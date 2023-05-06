package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, create Create) error
}
