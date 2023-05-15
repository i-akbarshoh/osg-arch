package user

import (
	"context"
)

type Repository interface {
	Create(context.Context, Create) error
	Get(context.Context, string) (Get, error)
	List(context.Context) (List, error)
	Update(context.Context, Update) error
	Delete(context.Context, Delete) error
}
