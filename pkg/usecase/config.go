package usecase

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/spf13/pflag"
)

type Config interface {
	Init(ctx context.Context, flags *pflag.FlagSet, args []string) error
	SetPath(ctx context.Context, flags *pflag.FlagSet, args []string) error
}

type ConfigUsecase struct {
	Repositories *domain.Repositories
}

func (u *ConfigUsecase) Init(ctx context.Context, flags *pflag.FlagSet, args []string) error {
	u.Repositories.ConfigRepository.Initialize()
	return nil
}

func (u *ConfigUsecase) SetPath(ctx context.Context, flags *pflag.FlagSet, args []string) error {
	str, err := flags.GetString("path")
	if err != nil {
		return err
	}
	if err := u.Repositories.ConfigRepository.SetWorkindDirectory(str); err != nil {
		return err
	}

	return nil
}

func NewConfigUsecase(r *domain.Repositories) Config {
	return &ConfigUsecase{
		Repositories: r,
	}
}
