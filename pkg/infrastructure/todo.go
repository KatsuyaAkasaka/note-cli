package infrastructure

import (
	"context"
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type todoRepository struct {
	Config *config.Config
}

var errNotFound = fmt.Errorf("not found")

func (r *todoRepository) Get(ctx context.Context, params *todo.GetParams) (*todo.Todo, error) {
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	contents, err := todoClient.ReadAll()
	if err != nil {
		return nil, err
	}
	todos := todo.FormatMD.ParseAll(contents).FilterBy(
		func(t *todo.Todo) bool {
			return t.ID == params.ID
		},
	)
	if len(todos) < 1 {
		return nil, errNotFound
	}

	return todos[0], nil
}

func (r *todoRepository) Create(ctx context.Context, t *todo.Todo) (*todo.Todo, error) {
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	if err := todoClient.AppendLine(t.ToLine(todo.FileTypeMarkdown)); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) SetDone(ctx context.Context, params *todo.SetDoneParams) (*todo.Todo, error) {
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	contents, err := todoClient.ReadAll()
	if err != nil {
		return nil, err
	}
	todos := todo.FormatMD.ParseAll(contents)
	filteredTodos := todos.FilterBy(
		func(t *todo.Todo) bool {
			return t.ID == params.ID
		},
	)
	if len(filteredTodos) < 1 {
		return nil, errNotFound
	}
	targetTodo := filteredTodos[0]
	targetTodo.Done = params.Done
	replacedTodos := todos.Replace(targetTodo)

	if err := todoClient.ReplaceAll(replacedTodos.ToLine(todo.FileTypeMarkdown)); err != nil {
		return nil, err
	}

	return targetTodo, nil
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
	todoClient := io.NewClient(r.Config.General.WorkingDirectory, r.Config.Todo.FileName, todo.FileTypeMarkdown.String())
	contents, err := todoClient.ReadAll()
	if err != nil {
		return nil, err
	}
	todos := todo.FormatMD.ParseAll(contents)
	deleteTodos := todos.FilterBy(
		func(t *todo.Todo) bool {
			return t.ID == params.ID
		},
	)
	if len(deleteTodos) < 1 {
		return nil, errNotFound
	}

	filteredTodos := todos.FilterBy(
		func(t *todo.Todo) bool {
			return t.ID != params.ID
		},
	)

	if err := todoClient.ReplaceAll(filteredTodos.ToLine(todo.FileTypeMarkdown)); err != nil {
		return nil, err
	}

	return deleteTodos[0], nil
}

func NewTodoRepository(c *config.Config) todo.Repository {
	return &todoRepository{
		Config: c,
	}
}
