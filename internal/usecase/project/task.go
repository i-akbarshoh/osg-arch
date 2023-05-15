package project

import (
	"context"
	"github.com/i-akbarshoh/osg-arch/internal/service/task"
	"github.com/i-akbarshoh/osg-arch/internal/entity"
)

func (u *UseCase) CreateTask(ctx context.Context, c entity.Task) (int, error) {
	return u.t.CreateTask(ctx, task.Create{
		Title: c.Title,
		Description: c.Description,
		StartAt: c.StartAt,
		FinishAt: c.FinishAt,
		ProjectID: c.ProjectID,
		UserID: c.ProgrammerID,
	})
}