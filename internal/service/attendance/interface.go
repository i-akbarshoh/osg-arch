package attendance

import "context"

type Repository interface{
	Create(context.Context, Create) error
	List(context.Context, string, string) (List, error)
}