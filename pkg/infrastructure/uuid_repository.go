package infrastructure

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/uuid"
	libuuid "github.com/google/uuid"
)

type uuidRepository struct{}

func (r *uuidRepository) Gen() string {
	return libuuid.NewString()
}

func NewUUIDRepository() uuid.Repository {
	return &uuidRepository{}
}
