package adapter

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			if err := a.Usecase.Init(ctx, flags, &usecase.InitParams{
				Args: args,
			}); err != nil {
				return err
			}
			return nil
		},
	}
	return c.ToCobraCommand()
}

func (a *Config) Reset() *cobra.Command {
	c := &Command{
		Command: "reset",
		Desc:    "reset config",
		Option:  a.Option,
		Aliases: []string{"ini"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			if err := a.Usecase.Reset(ctx, flags, &usecase.ResetParams{
				Args: args,
			}); err != nil {
				return err
			}
			return nil
		},
	}
	return c.ToCobraCommand()
}

func (a *Config) SetWorkingDirectory() *cobra.Command {
	type params struct {
		path string
	}
	p := &params{}

	c := &Command{
		Command: "set-path",
		Desc:    "set store path",
		Option:  a.Option,
		Aliases: []string{"sp"},
		Exec: func(ctx context.Context, flags *pflag.FlagSet, args []string) error {
			if err := a.Usecase.SetPath(ctx, flags, &usecase.SetPathParams{
				Path: p.path,
				Args: args,
			}); err != nil {
				return err
			}
			return nil
		},
		SetFlags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().StringVarP(
				&p.path,
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
