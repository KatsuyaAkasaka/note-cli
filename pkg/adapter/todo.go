package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Todo struct {
	Usecase usecase.Todo
	Config  *config.Config
}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Exec:    a.Usecase.Add,
		Args:    cobra.ExactArgs(1),
		Option:  NewOption().Apply(KindTodo, a.Config),
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

func NewTodoAdatper(r *domain.Repositories) *Todo {
	config, err := r.ConfigRepository.Get(&config.ConfigGetParams{
		Overwrite:     true,
		NotFoundAsErr: false,
	})
	if err != nil {
		return nil
	}
	return &Todo{
		Usecase: usecase.NewTodoUsecase(r),
		Config:  config,
	}
}
