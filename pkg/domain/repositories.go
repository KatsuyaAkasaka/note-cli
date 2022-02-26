package domain

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/setting"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
)

type Repositories struct {
	TodoRepository   todo.Repository
	ConfigRepository config.Repository
	Setting          setting.Setting
}

func NewRepository(
	todoRepository todo.Repository,
	configRepository config.Repository,
	setting setting.Setting,
) *Repositories {
	return &Repositories{
		TodoRepository:   todoRepository,
		ConfigRepository: configRepository,
		Setting:          setting,
	}
}
