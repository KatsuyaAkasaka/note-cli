package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Config struct {
	Usecase usecase.Config
	Config  *config.Config
}

func (a *Config) Initialize() *cobra.Command {
	c := &Command{
		Command: "init",
		Desc:    "initialize config",
		Option:  NewOption().Apply(KindConfig, a.Config),
		Aliases: []string{"ini"},
		Exec:    a.Usecase.Init,
	}
	return c.ToCobraCommand()
}

func (a *Config) SetWorkingDirectory() *cobra.Command {
	c := &Command{
		Command: "set-path",
		Desc:    "set store path",
		Option:  NewOption().Apply(KindConfig, a.Config),
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

func NewConfigAdatper(r *domain.Repositories) *Config {
	config, err := r.ConfigRepository.Get(&config.ConfigGetParams{
		Overwrite:     true,
		NotFoundAsErr: false,
	})
	if err != nil {
		return nil
	}
	return &Config{
		Usecase: usecase.NewConfigUsecase(r),
		Config:  config,
	}
}
