package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
	"github.com/i-akbarshoh/osg-arch/internal/service/project"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func New(db *bun.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, p project.Create) (id int, err error) {
	body := entity.Project{
		Name: p.Name,
		Attachment: p.Attachment,
		TeamLeadID: p.TeamLeadID,
		Status: p.Status,
	}
	_, err = r.db.NewInsert().Model(&body).Returning("id").Exec(ctx, &id)
	return
}

func (r *repository) List(ctx context.Context) (entity.List, error) {
	list := entity.List{}
	p := entity.Project{}
	_, err := r.db.NewSelect().Model(&p).Exec(ctx, &list.Projects)
	if err != nil {
		return entity.List{}, err
	}
	list.Count = len(list.Projects)
	return list, nil
}

func (r *repository) UpdateStatus(ctx context.Context, u project.Update) (err error) {
	m := entity.Project{
		ID: u.ID,
		Status: u.Status,
	}
	_, err = r.db.NewUpdate().OmitZero().Model(&m).WherePK().Exec(ctx)
	return
}

func (r *repository) Delete(ctx context.Context, id int) (err error) {
	p := entity.Project{
		ID: id,
	}
	_, err = r.db.NewDelete().Model(&p).WherePK().Exec(ctx)
	return
}