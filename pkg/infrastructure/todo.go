package infrastructure

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type todoRepository struct {
	Config *config.Config
}

func (r *todoRepository) Create(ctx context.Context, t *todo.Todo) (*todo.Todo, error) {
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	if err := io.AppendLine(todoClient, t.ToLine(todo.FileTypeMarkdown)); err != nil {
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
func (r *todoRepository) List(ctx context.Context, params *todo.ListParams) (todo.Todos, error) {
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	contents, err := todoClient.ReadAll()
	if err != nil {
		return nil, err
	}
	todos := todo.FormatMD.ParseAll(contents)
	return todos.FilterBy(
		func(t *todo.Todo) bool {
			return t.Content != ""
		},
	), nil
}
func (r *todoRepository) Delete(ctx context.Context, params *todo.DeleteParams) (*todo.Todo, error) {
	return nil, nil
}

func NewTodoRepository(c *config.Config) todo.Repository {
	return &todoRepository{
		Config: c,
	}
}
