package user

import (
	"context"

	u "github.com/i-akbarshoh/osg-arch/internal/service/user"
	"github.com/uptrace/bun"
)

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, create u.Create) error {
	if _, err := r.db.NewInsert().Model(&create).ModelTableExpr("users").Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(ctx context.Context, id string) (get u.Get, err error) {
	_, err = r.db.NewSelect().Table("users").Exec(ctx, &get)
	return
}

func (r *repository) List(ctx context.Context) (list u.List, err error) {
	l := make([]struct {
		ID        string `json:"id" bun:"id"`
		FullName  string `json:"full_name" bun:"full_name"`
		Password  string `bun:"password"`
		Avatar    string `json:"avatar" bun:"avatar"`
		Role      string `json:"role" bun:"role"`
		BirthDate string `json:"birth_date" bun:"birth_date"`
		Phone     string `json:"phone" bun:"phone"`
		Position  string `json:"position" bun:"position"`
	}, 0)
	count, err := r.db.NewSelect().Table("users").ScanAndCount(ctx, &l)
	list.Count = count
	for _, v := range l {
		list.L = append(list.L, u.Get{
			ID:        v.ID,
			FullName:  v.FullName,
			Avatar:    v.Avatar,
			Role:      v.Role,
			BirthDate: v.BirthDate,
			Phone:     v.Phone,
			Position:  v.Position,
		})
	}

	return
}
