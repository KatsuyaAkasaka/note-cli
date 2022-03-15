package usecase

import (
	"context"
	"strings"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/spf13/pflag"
)

type Todo interface {
	Add(ctx context.Context, flags *pflag.FlagSet, params *AddParams) (*todo.Todo, error)
	List(ctx context.Context, flags *pflag.FlagSet, params *ListParams) (todo.Todos, error)
	Switch(ctx context.Context, flags *pflag.FlagSet, params *SwitchParams) (*todo.Todo, error)
	Delete(ctx context.Context, flags *pflag.FlagSet, params *DeleteParams) (*todo.Todo, error)
}

type TodoUsecase struct {
	Repositories *domain.Repositories
}

type AddParams struct {
	Done bool
	Args []string
}

func (u *TodoUsecase) Add(ctx context.Context, flags *pflag.FlagSet, params *AddParams) (*todo.Todo, error) {
	t := &todo.Todo{
		ID:      u.Repositories.UUID.Gen(),
		Content: strings.Join(params.Args, " "),
		Done:    params.Done,
	}
	t, err := u.Repositories.TodoRepository.Create(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

type ListParams struct {
	Args []string
}

func (u *TodoUsecase) List(ctx context.Context, flags *pflag.FlagSet, params *ListParams) (todo.Todos, error) {
	todos, err := u.Repositories.TodoRepository.List(ctx, &todo.ListParams{})
	if err != nil {
		return nil, err
	}
	return todos, nil
}

type SwitchParams struct {
	ID   string
	Args []string
}

func (u *TodoUsecase) Switch(ctx context.Context, flags *pflag.FlagSet, params *SwitchParams) (*todo.Todo, error) {
	t, err := u.Repositories.TodoRepository.Get(ctx, &todo.GetParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, err
	}
	t, err = u.Repositories.TodoRepository.SetDone(ctx, &todo.SetDoneParams{
		ID:   params.ID,
		Done: !t.Done,
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

type DeleteParams struct {
	ID   string
	Args []string
}

func (u *TodoUsecase) Delete(ctx context.Context, flags *pflag.FlagSet, params *DeleteParams) (*todo.Todo, error) {
	t, err := u.Repositories.TodoRepository.Get(ctx, &todo.GetParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, err
	}
	t, err = u.Repositories.TodoRepository.Delete(ctx, &todo.DeleteParams{
		ID: params.ID,
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func NewTodoUsecase(r *domain.Repositories) Todo {
	return &TodoUsecase{
		Repositories: r,
	}
}
