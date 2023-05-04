package user

import "github.com/uptrace/bun"

type repository struct {
	*bun.DB
}

func NewRepository(db *bun.DB) *repository {
	return &repository{db}
}
