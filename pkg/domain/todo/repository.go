package todo

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, t *Todo) (*Todo, error)
	Update(ctx context.Context, t *Todo) (*Todo, error)
	SetDone(ctx context.Context, params *SetDoneParams) (*Todo, error)
	List(ctx context.Context, params *ListParams) (Todos, error)
	Delete(ctx context.Context, params *DeleteParams) (*Todo, error)
}

type DeleteParams struct {
	ID string
}

type SetDoneParams struct {
	ID string
}

type ListParams struct {
}
