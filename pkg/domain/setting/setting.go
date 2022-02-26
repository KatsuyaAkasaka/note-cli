package setting

import "github.com/KatsuyaAkasaka/nt/pkg/domain/config"

type Setting interface {
	Get(params *GetParams) (*config.Config, error)
}

type GetParams struct {
	Type string
}
