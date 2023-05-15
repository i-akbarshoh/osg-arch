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
	res, err := r.db.NewInsert().Returning("id").Exec(ctx, c)
	if err != nil {
		return
	}
	ID, err := res.LastInsertId()
	id = int(ID)
	return
}