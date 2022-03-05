package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Todo struct {
	Usecase usecase.Todo
	Option  *Option
}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Exec:    a.Usecase.Add,
		Args:    cobra.ExactArgs(1),
		Option:  a.Option,
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().BoolP(
				"done",
				"d",
				false,
				"set todo is done or not",
			)
		},
	}
	return c.ToCobraCommand()
}

func NewTodoAdatper(r *domain.Repositories, c *config.Config) *Todo {
	return &Todo{
		Usecase: usecase.NewTodoUsecase(r),
		Option:  NewOption().Apply(KindTodo, c),
	}
}
