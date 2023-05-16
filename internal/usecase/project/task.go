package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/task"
)

func (u *UseCase) CreateTask(ctx context.Context, c entity.Task) (int, error) {
	return u.t.CreateTask(ctx, task.Create{
		Title:       c.Title,
		Description: c.Description,
		StartAt:     c.StartAt,
		FinishAt:    c.FinishAt,
		ProjectID:   c.ProjectID,
		UserID:      c.ProgrammerID,
	})
}

func (u *UseCase) UpdateTask(ctx context.Context, up entity.Task) error {
	return u.t.Update(ctx, task.Update{
		ID:     up.ID,
		Status: up.Status,
	})
}

func (u *UseCase) GetTask(ctx context.Context, id int) (entity.Task, error) {
	g, err := u.t.Get(ctx, id)
	if err != nil {
		return entity.Task{}, err
	}
	res := entity.Task{
		ID:          g.ID,
		Title:       g.Title,
		Description: g.Description,
		StartAt:     g.StartAt,
		FinishAt:    g.FinishAt,
		StartedAt:   g.StartedAt,
		FinishedAt:  g.FinishedAt,
		Status:      g.UserID,
		ProjectID:   g.ProjectID,
		Attachment:  g.Attachment,
	}

	return res, nil
}

func (u *UseCase) ListTask(ctx context.Context) (l entity.ListTasks, err error) {
	r, err := u.t.ListTasks(ctx)
	if err != nil {
		return
	}

	for _, g := range r.List {
		l.List = append(l.List, entity.Task{
			ID:          g.ID,
			Title:       g.Title,
			Description: g.Description,
			StartAt:     g.StartAt,
			FinishAt:    g.FinishAt,
			StartedAt:   g.StartedAt,
			FinishedAt:  g.FinishedAt,
			Status:      g.UserID,
			ProjectID:   g.ProjectID,
			Attachment:  g.Attachment,
		})
	}

	l.Count = r.Count
	return
}

func (u *UseCase) DeleteTask(ctx context.Context, id int) error {
	return u.t.DeleteTask(ctx, id)
}