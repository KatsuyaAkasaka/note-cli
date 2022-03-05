package domain

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/uuid"
)

type Repositories struct {
	TodoRepository   todo.Repository
	ConfigRepository config.Repository
	UUID             uuid.Repository
}

func NewRepository(
	todoRepository todo.Repository,
	configRepository config.Repository,
	uuidRepository uuid.Repository,
) *Repositories {
	return &Repositories{
		TodoRepository:   todoRepository,
		ConfigRepository: configRepository,
		UUID:             uuidRepository,
	}
}
