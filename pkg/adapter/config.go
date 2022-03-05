package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Config struct {
	Usecase usecase.Config
	Option  *Option
}

func (a *Config) Initialize() *cobra.Command {
	c := &Command{
		Command: "init",
		Desc:    "initialize config",
		Option:  a.Option,
		Aliases: []string{"ini"},
		Exec:    a.Usecase.Init,
	}
	return c.ToCobraCommand()
}

func (a *Config) Reset() *cobra.Command {
	c := &Command{
		Command: "reset",
		Desc:    "reset config",
		Option:  a.Option,
		Aliases: []string{"ini"},
		Exec:    a.Usecase.Reset,
	}
	return c.ToCobraCommand()
}

func (a *Config) SetWorkingDirectory() *cobra.Command {
	c := &Command{
		Command: "set-path",
		Desc:    "set store path",
		Option:  a.Option,
		Aliases: []string{"sp"},
		Exec:    a.Usecase.SetPath,
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().StringP(
				"path",
				"p",
				"",
				"set working directory path to store where you want",
			)
		},
	}
	return c.ToCobraCommand()
}

func NewConfigAdatper(r *domain.Repositories, c *config.Config) *Config {
	return &Config{
		Usecase: usecase.NewConfigUsecase(r),
		Option:  NewOption().Apply(KindTodo, c),
	}
}
