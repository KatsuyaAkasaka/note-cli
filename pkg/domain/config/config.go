package config

import "fmt"

type Config struct {
	Note_cli struct {
		Timeout           int
		Working_directory string
		Todo              struct {
			Filename string
		}
	}
}

func (c *Config) TodoPath() string {
	return fmt.Sprintf(c.Note_cli.Working_directory + "/" + c.Note_cli.Todo.Filename)
}
