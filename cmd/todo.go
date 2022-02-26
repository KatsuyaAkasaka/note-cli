package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/adapter"
	"github.com/spf13/cobra"
)

// todo todo command
func setTodo() *cobra.Command {
	var todoCmd = &cobra.Command{
		Use:   "todo",
		Short: "todo list command",
		Long:  ``,
	}
	a := adapter.NewTodoAdatper()
	todoCmd.AddCommand(
		a.Add(),
	)
	return todoCmd
}
