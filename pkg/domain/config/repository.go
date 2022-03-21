package config

type Repository interface {
	Get(params *GetParams) (*Config, error)
	Initialize() (*Config, error)
	Reset() error
	SetWorkindDirectory(path string) error
}

type GetParams struct {
	Overwrite     bool
	NotFoundAsErr bool
}
