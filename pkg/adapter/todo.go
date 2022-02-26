package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Todo struct {
	Usecase usecase.Todo
}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Exec:    a.Usecase.Add,
	}
	return c.Convert().ToCobraCommand()
}

func NewTodoAdatper(r *domain.Repositories) *Todo {
	return &Todo{
		Usecase: usecase.NewTodoUsecase(r),
	}
}
