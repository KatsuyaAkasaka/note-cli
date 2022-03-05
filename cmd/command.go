package cmd

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "nt",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

type Commands interface {
	Cmd() *cobra.Command
}

func Cmd() *cobra.Command {
	config, err := infrastructure.NewConfigRepository().Get(&config.ConfigGetParams{
		Overwrite:     true,
		NotFoundAsErr: false,
	})
	if err != nil {
		return nil
	}
	rs := domain.NewRepository(
		infrastructure.NewTodoRepository(config),
		infrastructure.NewConfigRepository(),
		infrastructure.NewUUIDRepository(),
	)
	commands := []Commands{
		NewTodoCommand(rs, config),
		NewConfigCommand(rs, config),
	}
	for i := range commands {
		cmd.AddCommand(commands[i].Cmd())
	}
	return cmd
}
