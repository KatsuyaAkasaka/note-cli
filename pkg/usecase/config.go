package usecase

import (
	"context"
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
)

type Config interface {
	Init(ctx context.Context, args []string) error
}

type ConfigUsecase struct {
	Repositories *domain.Repositories
}

func (u *ConfigUsecase) Init(ctx context.Context, args []string) error {
	fmt.Println("hello world config init")
	return nil
}

func NewConfigUsecase(r *domain.Repositories) Config {
	return &ConfigUsecase{
		Repositories: r,
	}
}
