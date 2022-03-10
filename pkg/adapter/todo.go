package adapter

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/adapter/marshaler"
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Todo struct {
	Usecase usecase.Todo
	Option  *Option
}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Aliases: []string{"a"},
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

func (a *Todo) List() *cobra.Command {
	c := &Command{
		Command: "list",
		Desc:    "list todos",
		Aliases: []string{"l", "ls"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			todos, err := a.Usecase.List(ctx, flags, args)
			withID, err := flags.GetBool("id")
			if err != nil {
				return err
			}
			output := marshaler.TodosToOutput(todos, withID)
			Outputs(output)
			return err
		},
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().BoolP(
				"id",
				"i",
				false,
				"If true, visible id",
			)
		},
		Option: a.Option,
	}
	return c.ToCobraCommand()
}

func (a *Todo) Switch() *cobra.Command {
	c := &Command{
		Command: "switch",
		Desc:    "switch todo done",
		Aliases: []string{"s", "sw"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			_, err := a.Usecase.Switch(ctx, flags, args)
			return err
		},
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().StringP(
				"id",
				"i",
				"",
				"switch done flag id",
			)
		},
		Option: a.Option,
	}
	return c.ToCobraCommand()
}

func NewTodoAdatper(r *domain.Repositories, c *config.Config) *Todo {
	return &Todo{
		Usecase: usecase.NewTodoUsecase(r),
		Option:  NewOption().Apply(KindTodo, c),
	}
}
