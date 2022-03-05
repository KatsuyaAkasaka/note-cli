package io

import (
	"fmt"
	"os"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/viper"
)

type ioIns struct {
	viper *viper.Viper
}

var (
	configIns     *ioIns
	predifinedIns *ioIns
)

func ConfigIo() *ioIns {
	if configIns == nil {
		v := viper.New()

		v.SetConfigName("note-cli")
		v.SetConfigType("yaml")
		v.AddConfigPath("$HOME/.config")

		configIns = toIO(v)
	}
	return configIns
}

func PredefinedIo() *ioIns {
	if predifinedIns == nil {
		v := viper.New()

		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("config")

		predifinedIns = toIO(v)
	}
	return predifinedIns
}

func toIO(v *viper.Viper) *ioIns {
	return &ioIns{
		viper: v,
	}
}

// Load set config data into ioIns
func (i *ioIns) Load() error {
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

// GetConfig get config in ioIns
func (i *ioIns) GetConfig() (*config.Config, error) {
	c := &config.Config{}
	if err := i.Load(); err != nil {
		return nil, err
	}

	if err := i.viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("config err: %w", err)
	}
	return c, nil
}

// Write
func (i *ioIns) Write() error {
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
func (i *ioIns) Create() error {
	if err := i.viper.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}

func (i *ioIns) WriteOrCreate() error {
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

func (i *ioIns) Set(key string, value interface{}) {
	i.viper.Set(key, value)
}

func (i *ioIns) CopyConfigTo(to *ioIns) error {
	c, err := i.GetConfig()
	if err != nil {
		return err
	}

	to.Set("general.timeout", c.General.Timeout)
	to.Set("general.working_directory", c.General.WorkingDirectory)
	to.Set("todo.file_name", c.Todo.FileName)

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

func Append(path, line string) error {
	test_file_append, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer test_file_append.Close()
	if _, err := fmt.Fprintln(test_file_append, line); err != nil {
		return err
	}
	return nil
}
