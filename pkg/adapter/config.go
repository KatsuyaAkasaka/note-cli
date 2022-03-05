package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/setting"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Config struct {
	Usecase usecase.Config
	Setting setting.Setting
}

func (a *Config) Initialize() *cobra.Command {
	c := &Command{
		Command: "init",
		Desc:    "initialize config",
		Aliases: []string{"ini"},
		Timeout: 3,
		Exec:    a.Usecase.Init,
	}
	return c.ToCobraCommand()
}

func (a *Config) SetWorkingDirectory() *cobra.Command {
	c := &Command{
		Command: "set-path",
		Desc:    "set store path",
		Aliases: []string{"sp"},
		Exec:    a.Usecase.SetPath,
		SetFlags: func(cmd *cobra.Command) {
			cmd.Flags().StringP(
				"path",
				"p",
				"",
				"set working directory path to store datas where you want",
			)
		},
	}
	setting, err := a.Setting.Get(&setting.GetParams{})
	if err != nil {
		return nil
	}
	return c.Apply(setting).ToCobraCommand()
}

func NewConfigAdatper(r *domain.Repositories) *Config {
	return &Config{
		Usecase: usecase.NewConfigUsecase(r),
		Setting: r.Setting,
	}
}
