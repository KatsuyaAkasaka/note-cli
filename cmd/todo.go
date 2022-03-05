package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/adapter"
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/cobra"
)

type TodoCommand struct {
	Repositories *domain.Repositories
	Config       *config.Config
}

// Cmd todo command
func (c *TodoCommand) Cmd() *cobra.Command {
	var todoCmd = &cobra.Command{
		Use:   "todo",
		Short: "todo management command",
		Long:  ``,
	}
	a := adapter.NewTodoAdatper(c.Repositories, c.Config)
	todoCmd.AddCommand(
		a.Add(),
		a.List(),
	)
	return todoCmd
}

func NewTodoCommand(r *domain.Repositories, c *config.Config) Commands {
	return &TodoCommand{
		Repositories: r,
		Config:       c,
	}
}
