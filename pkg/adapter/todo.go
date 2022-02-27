package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/setting"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Todo struct {
	Usecase usecase.Todo
	Setting setting.Setting
}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Exec:    a.Usecase.Add,
		Args:    cobra.ExactArgs(1),
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().BoolP(
				"done",
				"d",
				false,
				"set todo is done or not",
			)
		},
	}
	setting, err := a.Setting.Get(&setting.GetParams{})
	if err != nil {
		return nil
	}
	return c.Apply(setting).ToCobraCommand()
}

func NewTodoAdatper(r *domain.Repositories) *Todo {
	return &Todo{
		Usecase: usecase.NewTodoUsecase(r),
		Setting: r.Setting,
	}
}
