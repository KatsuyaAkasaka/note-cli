package todo

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, params *GetParams) (*Todo, error)
	Create(ctx context.Context, t *Todo) (*Todo, error)
	SetDone(ctx context.Context, params *SetDoneParams) (*Todo, error)
	List(ctx context.Context, params *ListParams) (Todos, error)
	Delete(ctx context.Context, params *DeleteParams) (*Todo, error)
}

type GetParams struct {
	ID string
}

type DeleteParams struct {
	ID string
}

type SetDoneParams struct {
	ID   string
	Done bool
}

type ListParams struct{}
