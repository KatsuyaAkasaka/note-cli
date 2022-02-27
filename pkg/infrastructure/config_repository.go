package infrastructure

import (
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type configRepository struct{}

const errPrefix = "config err:"

func (r *configRepository) Initialize() (*config.Config, error) {
	configIo := io.ConfigIo()

	c, err := io.Get(configIo)
	if err != nil {
		if io.IsErrNotFound(err) {
			if err := io.Copy(io.PredifinedIo(), configIo); err != nil {
				return nil, fmt.Errorf("%s %w", errPrefix, err)
			}

			c, err = io.Get(configIo)
			if err != nil {
				return nil, fmt.Errorf("%s %w", errPrefix, err)
			}
		} else {
			return nil, fmt.Errorf("%s %w", errPrefix, err)
		}
	}
	return c, nil
}
func (r *configRepository) Reset() error {
	configViper := io.ConfigIo()

	io.SetDefault(configViper)
	if err := io.Create(configViper); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func (r *configRepository) SetWorkindDirectory(path string) error {
	if err := io.Exists(path); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	confIO := io.ConfigIo()
	if err := io.Set(confIO, "note_cli.working_directory", path); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func NewConfigRepository() config.Repository {
	return &configRepository{}
}
