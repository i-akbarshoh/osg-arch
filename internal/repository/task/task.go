package task

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/service/task"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func New(db *bun.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, c task.Create) (id int, err error) {
	_, err = r.db.NewInsert().Model(&c).Returning("id").Exec(ctx, &id)
	if err != nil {
		return
	}
	return
}

func (r *repository) UpdateStatus(ctx context.Context,u task.Update) (err error) {
	_, err = r.db.NewUpdate().OmitZero().Model(&u).WherePK().Exec(ctx)
	return
}

func (r *repository) Get(ctx context.Context, id int) (g task.Get, err error) {
	_, err = r.db.NewSelect().Model(&task.Get{}).Where("id = ?", id).Exec(ctx, &g)
	return
}

func (r *repository) List(ctx context.Context) (l task.List, err error) {
	_, err = r.db.NewSelect().Model(&task.Get{}).Exec(ctx, &l.List)
	l.Count = len(l.List)
	return
}

func (r *repository) DeleteTask(ctx context.Context, id int) (err error) {
	_, err = r.db.NewDelete().Model(&task.Get{
		ID: id,
	}).WherePK().Exec(ctx)
	return
}