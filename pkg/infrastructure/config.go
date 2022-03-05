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
	ConfigClient := io.ConfigClient()
	PredefinedClient := io.PredefinedClient()

	dst, err := ConfigClient.GetConfig(true)
	if err == nil {
		return dst, nil
	}

	if !io.IsErrNotFound(err) {
		return nil, fmt.Errorf("%s %w", errPrefix, err)
	}

	if err := PredefinedClient.CopyConfigTo(ConfigClient); err != nil {
		return nil, fmt.Errorf("%s %w", errPrefix, err)
	}

	if err := ConfigClient.Create(); err != nil {
		if err != nil {
			return nil, fmt.Errorf("%s %w", errPrefix, err)
		}
	}
	return ConfigClient.GetConfigWithOverwriteDefault(true)
}
func (r *configRepository) Reset() error {
	ConfigClient := io.ConfigClient()
	PredefinedClient := io.PredefinedClient()
	if err := PredefinedClient.CopyConfigTo(ConfigClient); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func (r *configRepository) SetWorkindDirectory(path string) error {
	if err := io.Exists(path); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	confIO := io.ConfigClient()
	confIO.Set("general.working_directory", path)
	if err := confIO.Write(); err != nil {
		return fmt.Errorf("%s %w", errPrefix, err)
	}
	return nil
}

func (r *configRepository) Get(params *config.ConfigGetParams) (*config.Config, error) {
	ConfigClient := io.ConfigClient()
	dst := &config.Config{}
	var err error
	if params.Overwrite {
		dst, err = ConfigClient.GetConfigWithOverwriteDefault(params.NotFoundAsErr)
	} else {
		dst, err = ConfigClient.GetConfig(params.NotFoundAsErr)
	}
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func NewConfigRepository() config.Repository {
	return &configRepository{}
}
