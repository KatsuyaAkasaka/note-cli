package adapter

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/KatsuyaAkasaka/nt/pkg/usecase"
	"github.com/spf13/cobra"
)

type Config struct {
	Usecase usecase.Config
}

func (a *Config) Initialize() *cobra.Command {
	c := &Command{
		Command: "init",
		Desc:    "initialize config",
		Aliases: []string{"ini"},
		Exec:    a.Usecase.Init,
	}
	return c.Convert().ToCobraCommand()
}

func NewConfigAdatper(r *domain.Repositories) *Config {
	return &Config{
		Usecase: usecase.NewConfigUsecase(r),
	}
}
