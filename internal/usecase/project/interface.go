package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
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
}
