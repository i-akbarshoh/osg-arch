package comment

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/i-akbarshoh/osg-arch/internal/service/comment"
)

type repository struct {
	db *bun.DB
}

func New(db *bun.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context,c comment.Create) (id int, err error) {
	_, err = r.db.NewInsert().Model(&c).Returning("id").Exec(ctx, &id)
	return
}

func (r *repository) List(ctx context.Context, task_id int) (l comment.List, err error) {
	_, err = r.db.NewSelect().Model(&comment.Create{}).Where("task_id = ?", &task_id).Exec(ctx, &l.L)
	l.Count = len(l.L)
	return
}

func (r *repository) Delete(ctx context.Context, id int) (err error) {
	_, err = r.db.NewDelete().Model(&comment.Create{ID: id}).WherePK().Exec(ctx)
	return
}