package config

type Repository interface {
	Initialize() (*Config, error)
	Get(params *GetParams) (*Config, error)
	Reset() error
}

type GetParams struct {
	Type string
}
