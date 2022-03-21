package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/adapter"
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/cobra"
)

type ConfigCommand struct {
	Repositories *domain.Repositories
	Config       *config.Config
}

// Cmd config command.
func (c *ConfigCommand) Cmd() *cobra.Command {
	configCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "config",
		Short: "config command",
		Long:  ``,
	}
	a := adapter.NewConfigAdatper(c.Repositories, c.Config)
	configCmd.AddCommand(
		a.Initialize(),
		a.Reset(),
		a.SetWorkingDirectory(),
	)

	return configCmd
}

func NewConfigCommand(r *domain.Repositories, c *config.Config) Commands {
	return &ConfigCommand{
		Repositories: r,
		Config:       c,
	}
}
