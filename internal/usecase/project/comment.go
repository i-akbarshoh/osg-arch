package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/comment"
)

func (u *UseCase) CreateComment(ctx context.Context, c entity.Comment) (int, error) {
	return u.c.Create(ctx, comment.Create{
		Text: c.Text,
		TaskID: c.TaskID,
		ProgrammerID: c.ProgrammerID,
	})
}

func (u *UseCase) ListComments(ctx context.Context, task_id int) (comment.List, error) {
	return u.c.ListComments(ctx, task_id)
}

func (u *UseCase) DeleteComment(ctx context.Context, id int) error {
	return u.c.DeleteComment(ctx, id)
}