package project

import (
	"context"

	"github.com/i-akbarshoh/osg-arch/internal/entity"
)

type Repository interface{
	Create(ctx context.Context,p Create) (int, error)
	List(ctx context.Context) (entity.List, error)
	UpdateStatus(ctx context.Context, p Update) error
	Delete(ctx context.Context, id int) error
}