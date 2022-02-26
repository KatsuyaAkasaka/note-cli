package domain

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
)

type Repositories struct {
	TodoRepository   *todo.Repository
	ConfigRepository *config.Repository
}

func NewRepository(
	todoRepository todo.Repository,
	configRepository config.Repository,
) *Repositories {
	return &Repositories{
		TodoRepository:   &todoRepository,
		ConfigRepository: &configRepository,
	}
}
