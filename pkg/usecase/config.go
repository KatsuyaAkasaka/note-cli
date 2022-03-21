package usecase

import (
	"context"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/spf13/pflag"
)

type Config interface {
	Init(ctx context.Context, flags *pflag.FlagSet, params *InitParams) error
	Reset(ctx context.Context, flags *pflag.FlagSet, params *ResetParams) error
	SetPath(ctx context.Context, flags *pflag.FlagSet, params *SetPathParams) error
}

type ConfigUsecase struct {
	Repositories *domain.Repositories
}

type InitParams struct {
	Args []string
}

func (u *ConfigUsecase) Init(ctx context.Context, flags *pflag.FlagSet, params *InitParams) error {
	if _, err := u.Repositories.ConfigRepository.Initialize(); err != nil {
		return err
	}

	return nil
}

type ResetParams struct {
	Args []string
}

func (u *ConfigUsecase) Reset(ctx context.Context, flags *pflag.FlagSet, params *ResetParams) error {
	if err := u.Repositories.ConfigRepository.Reset(); err != nil {
		return err
	}

	return nil
}

type SetPathParams struct {
	Path string
	Args []string
}

func (u *ConfigUsecase) SetPath(ctx context.Context, flags *pflag.FlagSet, params *SetPathParams) error {
	if err := u.Repositories.ConfigRepository.SetWorkindDirectory(params.Path); err != nil {
		return err
	}

	return nil
}

func NewConfigUsecase(r *domain.Repositories) Config {
	return &ConfigUsecase{
		Repositories: r,
	}
}
