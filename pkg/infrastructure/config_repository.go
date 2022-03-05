package infrastructure

import (
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type configRepository struct{}

const errPrefix = "config err:"

// Initialize initialize config based on predefined config
func (r *configRepository) Initialize() (*config.Config, error) {
	configIO := io.ConfigIo()
	predefinedIO := io.PredefinedIo()

	dst, err := configIO.GetConfig()
	if err == nil {
		return dst, nil
	}

	if !io.IsErrNotFound(err) {
		return nil, fmt.Errorf("%s %w", errPrefix, err)
	}

	if err := predefinedIO.CopyConfigTo(configIO); err != nil {
		return nil, fmt.Errorf("%s %w", errPrefix, err)
	}

	if err := configIO.Create(); err != nil {
		if err != nil {
			return nil, fmt.Errorf("%s %w", errPrefix, err)
		}
	}
	return configIO.GetConfig()
}
func (r *configRepository) Reset() error {
	configIO := io.ConfigIo()
	predefinedIO := io.PredefinedIo()
	if err := predefinedIO.CopyConfigTo(configIO); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func (r *configRepository) SetWorkindDirectory(path string) error {
	if err := io.Exists(path); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	confIO := io.ConfigIo()
	confIO.Set("general.working_directory", path)
	if err := confIO.Write(); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func NewConfigRepository() config.Repository {
	return &configRepository{}
}
