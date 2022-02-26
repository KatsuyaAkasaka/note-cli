package config

type Repository interface {
	Initialize() (*Config, error)
	Reset() error
}
