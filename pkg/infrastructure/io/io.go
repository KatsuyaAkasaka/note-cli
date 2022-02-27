package io

import (
	"fmt"
	"os"

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

func Load(v *viper.Viper) error {
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func SetDefault(v *viper.Viper) {
	v.SetDefault("note_cli.timeout", 5)
}

func Set(v *viper.Viper, conf string, val interface{}) error {
	v.Set(conf, val)
	if err := Load(v); err != nil {
		return err
	}
	if err := v.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func Get(v *viper.Viper) (*config.Config, error) {
	c := &config.Config{}
	if err := Load(v); err != nil {
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
	if err := v.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config err: %w", err)
		} else {
			return &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	return nil
}

func Create(v *viper.Viper) error {
	if err := v.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}

func WriteOrCreate(v *viper.Viper) error {
	if err := Write(v); err != nil {
		if IsErrNotFound(err) {
			if err := Create(v); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Copy(from, to *viper.Viper) error {
	c, err := Get(from)
	if err != nil {
		return fmt.Errorf("config err: %w", err)
	}

	to.Set("note_cli.timeout", c.Note_cli.Timeout)

	if err := Create(to); err != nil {
		return err
	}
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
