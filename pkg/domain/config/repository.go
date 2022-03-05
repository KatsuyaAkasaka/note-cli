package config

type Repository interface {
	Get(params *ConfigGetParams) (*Config, error)
	Initialize() (*Config, error)
	Reset() error
	SetWorkindDirectory(path string) error
}

type ConfigGetParams struct {
	Overwrite     bool
	NotFoundAsErr bool
}
