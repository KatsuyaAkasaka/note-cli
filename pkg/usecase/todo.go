package usecase

import (
	"context"
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
	"github.com/spf13/pflag"
)

type Todo interface {
	Add(ctx context.Context, flags *pflag.FlagSet, args []string) error
}

type TodoUsecase struct {
	Repositories *domain.Repositories
}

func (u *TodoUsecase) Add(ctx context.Context, flags *pflag.FlagSet, args []string) error {
	fmt.Println("hello world")
	return nil
}

func NewTodoUsecase(r *domain.Repositories) Todo {
	return &TodoUsecase{
		Repositories: r,
	}
}
