package usecase

import (
	"context"
	"strings"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/spf13/pflag"
)

type Todo interface {
	Add(ctx context.Context, flags *pflag.FlagSet, args []string) error
	List(ctx context.Context, flags *pflag.FlagSet, args []string) (todo.Todos, error)
	Switch(ctx context.Context, flags *pflag.FlagSet, args []string) (*todo.Todo, error)
}

type TodoUsecase struct {
	Repositories *domain.Repositories
}

func (u *TodoUsecase) Add(ctx context.Context, flags *pflag.FlagSet, args []string) error {
	done, err := flags.GetBool("done")
	if err != nil {
		return err
	}
	t := &todo.Todo{
		ID:      u.Repositories.UUID.Gen(),
		Content: strings.Join(args, " "),
		Done:    done,
	}
	t, err = u.Repositories.TodoRepository.Create(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (u *TodoUsecase) List(ctx context.Context, flags *pflag.FlagSet, args []string) (todo.Todos, error) {
	todos, err := u.Repositories.TodoRepository.List(ctx, &todo.ListParams{})
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (u *TodoUsecase) Switch(ctx context.Context, flags *pflag.FlagSet, args []string) (*todo.Todo, error) {
	id, err := flags.GetString("id")
	if err != nil {
		return nil, err
	}
	t, err := u.Repositories.TodoRepository.Get(ctx, &todo.GetParams{
		ID: id,
	})
	t, err = u.Repositories.TodoRepository.SetDone(ctx, &todo.SetDoneParams{
		ID:   id,
		Done: !t.Done,
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
