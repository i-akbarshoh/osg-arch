package task

import "context"

type Repository interface {
	Create(context.Context, Create) (int, error)
	UpdateStatus(context.Context, Update) error
	Get(context.Context, int) (Get, error)
	List(context.Context) (List, error)
	DeleteTask(context.Context, int) error
}