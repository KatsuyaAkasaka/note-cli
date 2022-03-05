package config

type (
	Config struct {
		General General `mapstructure:"general"`
		Todo    Todo    `mapstructure:"todo"`
	}
	General struct {
		WorkingDirectory string `mapstructure:"working_directory"`
	}
	Todo struct {
		FileName string `mapstructure:"file_name"`
		Timeout  int    `mapstructure:"timeout"`
	}
)

func (c *Config) Overwrite(to *Config) *Config {
	if to.General.WorkingDirectory != "" {
		c.General.WorkingDirectory = to.General.WorkingDirectory
	}
	if to.Todo.FileName != "" {
		c.Todo.FileName = to.Todo.FileName
	}
	if to.Todo.Timeout != 0 {
		c.Todo.Timeout = to.Todo.Timeout
	}
	return c
}
