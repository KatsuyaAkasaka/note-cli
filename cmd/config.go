package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/adapter"
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/spf13/cobra"
)

type ConfigCommand struct {
	Repositories *domain.Repositories
}

// Cmd config command
func (c *ConfigCommand) Cmd() *cobra.Command {
	var configCmd = &cobra.Command{
		Use:   "config",
		Short: "config command",
		Long:  ``,
	}
	a := adapter.NewConfigAdatper(c.Repositories)
	configCmd.AddCommand(
		a.Initialize(),
	)
	return configCmd
}

func NewConfigCommand(r *domain.Repositories) Commands {
	return &ConfigCommand{
		Repositories: r,
	}
}
