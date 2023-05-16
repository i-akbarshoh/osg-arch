package attendance

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/service/attendance"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func New(db *bun.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context,c attendance.Create) (err error) {
	_, err = r.db.NewInsert().Model(&c).Exec(ctx)
	return
}

func (r *repository) List(ctx context.Context, user_id string, t string) (list attendance.List,err error) {
	_, err = r.db.NewSelect().Column("type", "date").Table("attendance").Where("user_id = ? and type = ?", user_id, t).Exec(ctx, &list)
	return
}