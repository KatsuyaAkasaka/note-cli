package io

import (
	"fmt"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/viper"
)

func ConfigIo() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.note-cli")
	return v
}

func PredifinedIo() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")
	return v
}

func SetDefault(v *viper.Viper) {
	v.SetDefault("note_cli.todo.timeout", 5)
}

func Get(v *viper.Viper) (*config.Config, error) {
	c := &config.Config{}
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		} else {
			return nil, &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

func Write(v *viper.Viper) error {
	if err := v.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config err: %w", err)
		} else {
			return &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	return nil
}

func Copy(from, to *viper.Viper) error {
	c, err := Get(from)
	if err != nil {
		return fmt.Errorf("config err: %w", err)
	}

	to.Set("note_cli.todo.timeout", c.Note_cli.Todo.Timeout)

	if err := Write(to); err != nil {
		return err
	}
	return nil
}
