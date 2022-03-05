package config

import "fmt"

type (
	Config struct {
		General General `yaml:"general"`
		Todo    Todo    `yaml:"todo"`
	}
	General struct {
		Timeout          int    `yaml:"timeout"`
		WorkingDirectory string `yarml:"working_directory"`
	}
	Todo struct {
		FileName string
	}
)

// var (
// 	_defGeneral General = General{
// 		Timeout:          5,
// 		WorkingDirectory: "",
// 	}
// 	_defTodo Todo    = Todo{}
// 	def      *Config = &Config{
// 		General: _defGeneral,
// 		Todo:    _defTodo,
// 	}
// )

func (c *Config) TodoPath() string {
	return fmt.Sprintf(c.General.WorkingDirectory + "/" + c.Todo.FileName)
}

func (c *Config) Overwrite(to *Config) *Config {
	if c.General.Timeout == 0 {
		c.General.Timeout = to.General.Timeout
	}
	if c.General.WorkingDirectory == "" {
		c.General.WorkingDirectory = to.General.WorkingDirectory
	}
	if c.Todo.FileName == "" {
		c.Todo.FileName = to.Todo.FileName
	}
	return c
}
