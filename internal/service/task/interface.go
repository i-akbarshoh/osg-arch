package task

import "context"

type Repository interface {
	Create(context.Context, Create) (int, error)
}