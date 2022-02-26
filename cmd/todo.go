package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/adapter"
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/spf13/cobra"
)

type TodoCommand struct {
	Repositories *domain.Repositories
}

// Cmd todo command
func (c *TodoCommand) Cmd() *cobra.Command {
	var todoCmd = &cobra.Command{
		Use:   "todo",
		Short: "todo list command",
		Long:  ``,
	}
	a := adapter.NewTodoAdatper(c.Repositories)
	todoCmd.AddCommand(
		a.Add(),
	)
	return todoCmd
}

func NewTodoCommand(r *domain.Repositories) Commands {
	return &TodoCommand{
		Repositories: r,
	}
}
