package adapter

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type Todo struct{}

func (a *Todo) Add() *cobra.Command {
	c := &Command{
		Command: "add",
		Desc:    "add todo list",
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println("hello world")
			return nil
		},
	}
	return c.Convert().ToCobraCommand()
}

func NewTodoAdatper() *Todo {
	return &Todo{}
}
