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
	type params struct {
		done bool
	}
	p := &params{}

	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Aliases: []string{"a"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			if _, err := a.Usecase.Add(ctx, flags, &usecase.AddParams{
				Done: p.done,
				Args: args,
			}); err != nil {
				return err
			}
			return nil
		},
		Args:   cobra.ExactArgs(1),
		Option: a.Option,
		SetFlags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().BoolVarP(
				&p.done,
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
	type params struct {
		WithID bool
	}
	p := &params{}

	c := &Command{
		Command: "list",
		Desc:    "list todos",
		Aliases: []string{"l", "ls"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			todos, err := a.Usecase.List(ctx, flags, &usecase.ListParams{
				Args: args,
			})
			if err != nil {
				return err
			}

			outputStrs := marshaler.OutputTodos(todos, &marshaler.OutputTodosParams{
				WithID: p.WithID,
			})
			Outputs(outputStrs)
			return nil
		},
		SetFlags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().BoolVar(
				&p.WithID,
				"id",
				false,
				"If true, visible id",
			)
		},
		Option: a.Option,
	}
	return c.ToCobraCommand()
}

func (a *Todo) Switch() *cobra.Command {
	type params struct {
		id string
	}
	p := &params{}

	c := &Command{
		Command: "switch",
		Desc:    "switch todo done",
		Aliases: []string{"s", "sw"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			if _, err := a.Usecase.Switch(ctx, flags, &usecase.SwitchParams{
				ID:   p.id,
				Args: args,
			}); err != nil {
				return err
			}
			return nil
		},
		SetFlags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().StringVar(
				&p.id,
				"id",
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
