package infrastructure

import (
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/viper"
)

type configRepository struct{}

func configViper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.note-cli")
	return v
}

func predifinedViper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	return v
}

func get(v *viper.Viper) (*config.Config, error) {
	c := &config.Config{}
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

func SetDefault(v *viper.Viper) {
	v.SetDefault("note_cli.todo.timeout", 5)
}

func write(v *viper.Viper) error {
	if err := v.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}

func copy(from, to *viper.Viper) error {
	c, err := get(from)
	if err != nil {
		return fmt.Errorf("config err: %w", err)
	}

	to.Set("note_cli.todo.timeout", c.Note_cli.Todo.Timeout)

	if err := write(to); err != nil {
		return err
	}
	return nil
}

func (r *configRepository) Initialize() (*config.Config, error) {
	configViper := configViper()

	c, err := get(configViper)
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("config err: %w", err)
		} else {
			if err := copy(predifinedViper(), configViper); err != nil {
				return nil, fmt.Errorf("config err: %w", err)
			}

			c, err = get(configViper)
			if err != nil {
				return nil, fmt.Errorf("config err: %w", err)
			}
		}
	}
	return c, nil
}
func (r *configRepository) Get(params *config.GetParams) (*config.Config, error) {
	configViper := configViper()

	c, err := get(configViper)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *configRepository) Reset() error {
	configViper := configViper()

	SetDefault(configViper)
	if err := write(configViper); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}
func NewConfigRepository() config.Repository {
	return &configRepository{}
}
