package io

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/utils"
	"github.com/spf13/viper"
)

type client struct {
	viper    *viper.Viper
	FullPath string
	DirPath  string
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
	return &client{
		viper:    v,
		FullPath: path.Join(p, name) + "." + ext,
		DirPath:  p,
	}
}

func ConfigClient() *client {
	if configClient == nil {
		configClient = NewClient("$HOME/.config/note-cli", "config", "yaml")
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

// Load set config data into client
func (c *client) Load() error {
	if err := c.viper.ReadInConfig(); err != nil {
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
func (c *client) GetConfig(notFoundAsErr bool) (*config.Config, error) {
	conf := &config.Config{}
	if err := c.Load(); err != nil {
		if IsErrNotFound(err) && !notFoundAsErr {
			return conf, nil
		}
		return nil, err
	}

	if err := c.viper.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("config err: %w", err)
	}
	return conf, nil
}

// GetConfigWithOverwriteDefault get default config overrides config file
func (c *client) GetConfigWithOverwriteDefault(notFoundAsErr bool) (*config.Config, error) {
	conf, err := c.GetConfig(notFoundAsErr)
	if err != nil {
		return nil, err
	}

	defaultC, err := DefaultClient().GetConfig(notFoundAsErr)
	if err != nil {
		return nil, err
	}

	return defaultC.Overwrite(conf), nil
}

// Write write config file if file exists
func (c *client) Write() error {
	if err := c.viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config err: %w", err)
		} else {
			return &ErrNotFound{Err: fmt.Errorf("get err: not found")}
		}
	}
	return nil
}

// Create create config file if file does not exists
func (c *client) Create() error {
	if err := os.MkdirAll(utils.AbsolutePath(c.DirPath), 0755); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	if err := c.viper.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	return nil
}

// WriteOrCreate create config file and directory if file or directory does not exists
func (c *client) WriteOrCreate() error {
	if err := c.Write(); err != nil {
		if !IsErrNotFound(err) {
			return err
		}
		if err := c.Create(); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

func (c *client) CopyConfigTo(to *client) error {
	conf, err := c.GetConfig(true)
	if err != nil {
		return err
	}

	to.Set("general.working_directory", conf.General.WorkingDirectory)
	to.Set("todo.file_name", conf.Todo.FileName)
	to.Set("todo.timeout", conf.Todo.Timeout)

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

func (c *client) AppendLine(line string) error {
	fp, err := os.OpenFile(utils.AbsolutePath(c.FullPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fmt.Fprintln(fp, line); err != nil {
		return err
	}
	return nil
}

func (c *client) ReplaceAll(lines []string) error {
	fp, err := os.OpenFile(utils.AbsolutePath(c.FullPath), os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer fp.Close()
	inputText := strings.Join(lines, "\n")
	if _, err := fmt.Fprintln(fp, inputText); err != nil {
		return err
	}
	return nil
}

func (c *client) ReadAll() ([]string, error) {
	fp, err := os.Open(utils.AbsolutePath(c.FullPath))
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(data), "\n")
	filteredContents := []string{}
	for i := range contents {
		if contents[i] != "" {
			filteredContents = append(filteredContents, contents[i])
		}
	}
	return filteredContents, nil
}
