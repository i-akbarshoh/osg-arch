package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/project"
)

type UseCase struct {
	p Project
	t Task
}

func New(p Project, t Task) *UseCase {
	return &UseCase{p: p, t: t}
}

func (u *UseCase) Create(ctx context.Context,p entity.Project) (int, error) {
	return u.p.Create(ctx, project.Create{
		Name: p.Name,
		Status: p.Status,
		TeamLeadID: p.TeamLeadID,
		Attachment: p.Attachment,
	})
}

func (u *UseCase) List(ctx context.Context) (entity.List, error) {
	return u.p.List(ctx)
}

func (u *UseCase) Update(ctx context.Context, up entity.Project) error {
	return u.p.Update(ctx, project.Update{
		ID: up.ID,
		Status: up.Status,
	})
}

func (u *UseCase) Delete(ctx context.Context,p int) error {
	return u.p.Delete(ctx, p)
}