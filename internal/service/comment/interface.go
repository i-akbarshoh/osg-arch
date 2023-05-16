package comment

import "context"

type Repository interface{
	Create(context.Context, Create) (int, error)
	List(context.Context, int) (List, error)
	Delete(context.Context, int) error
}