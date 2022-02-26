package infrastructure

import (
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type configRepository struct{}

func (r *configRepository) Initialize() (*config.Config, error) {
	configIo := io.ConfigIo()

	c, err := io.Get(configIo)
	if err != nil {
		if io.IsErrNotFound(err) {
			return nil, fmt.Errorf("config err: %w", err)
		} else {
			if err := io.Copy(io.PredifinedIo(), configIo); err != nil {
				return nil, fmt.Errorf("config err: %w", err)
			}

			c, err = io.Get(configIo)
			if err != nil {
				return nil, fmt.Errorf("config err: %w", err)
			}
		}
	}
	return c, nil
}
func (r *configRepository) Reset() error {
	configViper := io.ConfigIo()

	io.SetDefault(configViper)
	if err := io.Write(configViper); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}
func NewConfigRepository() config.Repository {
	return &configRepository{}
}
