package infrastructure

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type todoRepository struct{}

func (r *todoRepository) Create(ctx context.Context, t *todo.Todo) (*todo.Todo, error) {
	configIO := io.ConfigIo()
	c, err := io.Get(configIO)
	if err != nil {
		return nil, err
	}
	if err := io.Append(c.TodoPath(), t.ToCheckMarkdown()); err != nil {
		return nil, err
	}
	return t, nil
}
func (r *todoRepository) Update(ctx context.Context, t *todo.Todo) (*todo.Todo, error) {
	return nil, nil
}
func (r *todoRepository) SetDone(ctx context.Context, params *todo.SetDoneParams) (*todo.Todo, error) {
	return nil, nil
}
func (r *todoRepository) List(ctx context.Context, params *todo.ListParams) (*todo.Todo, error) {
	return nil, nil
}
func (r *todoRepository) Delete(ctx context.Context, params *todo.DeleteParams) (*todo.Todo, error) {
	return nil, nil
}

func NewTodoRepository() todo.Repository {
	return &todoRepository{}
}
