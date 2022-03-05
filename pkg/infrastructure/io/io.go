package io

import (
	"fmt"
	"os"
	"path"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/viper"
)

type client struct {
	viper *viper.Viper
	Path  string
}

var (
	configClient     *client
	predefinedClient *client
	defaultClient    *client
)

func NewClient(p, name, ext string) *client {
	v := viper.New()

	v.AddConfigPath(p)
	v.SetConfigName(name)
	v.SetConfigType(ext)
	return toIO(v, path.Join(p, name)+"."+ext)
}

func ConfigClient() *client {
	if configClient == nil {
		configClient = NewClient("$HOME/.config", "note-cli", "yaml")
	}
	return configClient
}

func PredefinedClient() *client {
	if predefinedClient == nil {
		predefinedClient = NewClient("config", "src", "yaml")
	}
	return predefinedClient
}

func DefaultClient() *client {
	if defaultClient == nil {
		defaultClient = NewClient("config", "default", "yaml")
	}
	return defaultClient
}

func toIO(v *viper.Viper, path string) *client {
	return &client{
		viper: v,
		Path:  path,
	}
}

// Load set config data into client
func (i *client) Load() error {
	if err := i.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config err: %w", err)
		} else {
			return &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	return nil
}

func SetDefault(v *viper.Viper) {
	v.SetDefault("general.timeout", 5)
}

// GetConfig get config from config file
func (i *client) GetConfig(notFoundAsErr bool) (*config.Config, error) {
	c := &config.Config{}
	if err := i.Load(); err != nil {
		if IsErrNotFound(err) && !notFoundAsErr {
			return c, nil
		}
		return nil, err
	}

	if err := i.viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("config err: %w", err)
	}
	return c, nil
}

// GetConfigWithOverwriteDefault get default config overrides config file
func (i *client) GetConfigWithOverwriteDefault(notFoundAsErr bool) (*config.Config, error) {
	// defaultC := &config.Config{}
	// var err error
	c, err := i.GetConfig(notFoundAsErr)
	if err != nil {
		return nil, err
	}

	defaultC, err := DefaultClient().GetConfig(notFoundAsErr)
	if err != nil {
		return nil, err
	}

	return defaultC.Overwrite(c), nil
}

// Write write config file if file exists
func (i *client) Write() error {
	if err := i.viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config err: %w", err)
		} else {
			return &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	return nil
}

// Create create config file if file does not exists
func (i *client) Create() error {
	if err := i.viper.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}

func (i *client) WriteOrCreate() error {
	if err := i.Write(); err != nil {
		if !IsErrNotFound(err) {
			return err
		}
		if err := i.Create(); err != nil {
			return err
		}
	}
	return nil
}

func (i *client) Set(key string, value interface{}) {
	i.viper.Set(key, value)
}

func (i *client) CopyConfigTo(to *client) error {
	c, err := i.GetConfig(true)
	if err != nil {
		return err
	}

	to.Set("general.working_directory", c.General.WorkingDirectory)
	to.Set("todo.file_name", c.Todo.FileName)
	to.Set("todo.timeout", c.Todo.Timeout)

	return nil
}

func Exists(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory is not found")
		}
		return fmt.Errorf("path is invalid")
	}
	return nil
}

func AppendLine(target *client, line string) error {
	test_file_append, err := os.OpenFile(target.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer test_file_append.Close()
	if _, err := fmt.Fprintln(test_file_append, line); err != nil {
		return err
	}
	return nil
}
