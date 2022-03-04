package infrastructure

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	domainSetting "github.com/KatsuyaAkasaka/nt/pkg/domain/setting"
	"github.com/KatsuyaAkasaka/nt/pkg/infrastructure/io"
)

type setting struct{}

func (r *setting) Get(params *domainSetting.GetParams) (*config.Config, error) {
	dst := &config.Config{}
	var err error
	configIO := io.ConfigIo()
	if err := configIO.Load(); err != nil {
		if !io.IsErrNotFound(err) {
			return nil, err
		}
	}
	dst, err = configIO.GetConfig()
	if err != nil {
		if !io.IsErrNotFound(err) {
			return nil, err
		} else {
			dst = &config.Config{}
		}
	}
	return dst.OverwriteDefault(), nil
}

func NewSetting() domainSetting.Setting {
	return &setting{}
}
