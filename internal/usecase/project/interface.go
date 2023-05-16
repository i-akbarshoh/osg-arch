package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/comment"
	"github.com/i-akbarshoh/osg-arch/internal/service/project"
	"github.com/i-akbarshoh/osg-arch/internal/service/task"
)

type Project interface{
	Create(ctx context.Context,p project.Create) (int, error)
	List(context.Context) (entity.List, error)
	Update(context.Context, project.Update) error
	Delete(context.Context, int) error
}

type Task interface {
	CreateTask(context.Context, task.Create) (int, error)
	Update(context.Context, task.Update) error
	Get(context.Context, int) (task.Get, error)
	ListTasks(context.Context) (task.List, error)
	DeleteTask(context.Context, int) error
}

type Comment interface{
	Create(context.Context, comment.Create) (int, error)
	ListComments(context.Context, int) (comment.List, error)
	DeleteComment(context.Context, int) error
}
