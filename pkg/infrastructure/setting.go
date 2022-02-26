package infrastructure

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	domainSetting "github.com/KatsuyaAkasaka/nt/pkg/domain/setting"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type setting struct{}

func (r *setting) Get(params *domainSetting.GetParams) (*config.Config, error) {
	configIo := io.ConfigIo()

	c, err := io.Get(configIo)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewSetting() domainSetting.Setting {
	return &setting{}
}
