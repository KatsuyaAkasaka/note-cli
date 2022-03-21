package io

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/KatsuyaAkasaka/nt/pkg/domain/utils"
	"github.com/spf13/viper"
)

type Client struct {
	viper    *viper.Viper
	FullPath string
	DirPath  string
}

var (
	configClient     *Client
	predefinedClient *Client
	defaultClient    *Client
)

func NewClient(p, name, ext string) *Client {
	v := viper.New()

	v.AddConfigPath(p)
	v.SetConfigName(name)
	v.SetConfigType(ext)

	return &Client{
		viper:    v,
		FullPath: path.Join(p, name) + "." + ext,
		DirPath:  p,
	}
}

func ConfigClient() *Client {
	if configClient == nil {
		configClient = NewClient("$HOME/.config/note-cli", "config", "yaml")
	}

	return configClient
}

func PredefinedClient() *Client {
	if predefinedClient == nil {
		predefinedClient = NewClient("config", "src", "yaml")
	}

	return predefinedClient
}

func DefaultClient() *Client {
	if defaultClient == nil {
		defaultClient = NewClient("config", "default", "yaml")
	}

	return defaultClient
}

// Load set config data into Client.
func (c *Client) Load() error {
	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //nolint:errorlint
			return fmt.Errorf("config err: %w", err)
		}

		return &NotFoundError{Err: fmt.Errorf("get err: not found: %w", err)}
	}

	return nil
}

// GetConfig get config from config file.
func (c *Client) GetConfig(notFoundAsErr bool) (*config.Config, error) {
	conf := &config.Config{} //nolint:exhaustivestruct
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

// GetConfigWithOverwriteDefault get default config overrides config file.
func (c *Client) GetConfigWithOverwriteDefault(notFoundAsErr bool) (*config.Config, error) {
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

// Write write config file if file exists.
func (c *Client) Write() error {
	if err := c.viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //nolint:errorlint
			return fmt.Errorf("config err: %w", err)
		}

		return &NotFoundError{Err: fmt.Errorf("get err: not found: %w", err)}
	}

	return nil
}

// Create create config file if file does not exists.
func (c *Client) Create() error {
	var perm fs.FileMode = 0o755
	if err := os.MkdirAll(utils.AbsolutePath(c.DirPath), perm); err != nil {
		return fmt.Errorf("config err: %w", err)
	}
	if err := c.viper.SafeWriteConfig(); err != nil {
		return fmt.Errorf("config err: %w", err)
	}

	return nil
}

// WriteOrCreate create config file and directory if file or directory does not exists.
func (c *Client) WriteOrCreate() error {
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

func (c *Client) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

func (c *Client) CopyConfigTo(to *Client) error {
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
			return fmt.Errorf("directory is not found: %w", err)
		}

		return fmt.Errorf("path is invalid: %w", err)
	}

	return nil
}

func (c *Client) AppendLine(line string) error {
	var perm fs.FileMode = 0o755
	fp, err := os.OpenFile(utils.AbsolutePath(c.FullPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fmt.Fprintln(fp, line); err != nil {
		return err
	}

	return nil
}

func (c *Client) ReplaceAll(lines []string) error {
	var perm fs.FileMode = 0o755
	fp, err := os.OpenFile(utils.AbsolutePath(c.FullPath), os.O_RDWR|os.O_APPEND|os.O_TRUNC, perm)
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

func (c *Client) ReadAll() ([]string, error) {
	fp, err := os.Open(utils.AbsolutePath(c.FullPath))
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
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
