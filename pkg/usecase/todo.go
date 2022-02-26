package usecase

import (
	"context"
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain"
)

type Todo interface {
	Add(ctx context.Context, args []string) error
}

type TodoUsecase struct {
	Repositories *domain.Repositories
}

func (u *TodoUsecase) Add(ctx context.Context, args []string) error {
	fmt.Println("hello world")
	return nil
}

func NewTodoUsecase(r *domain.Repositories) Todo {
	return &TodoUsecase{
		Repositories: r,
	}
}
