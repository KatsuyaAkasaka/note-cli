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
	dst, err = configIO.GetConfig()
	if err != nil {
		if !io.IsErrNotFound(err) {
			return nil, err
		} else {
			dst = &config.Config{}
		}
	}
	c, err := r.Default()
	if err != nil {
		return nil, err
	}
	return dst.Overwrite(c), nil
}

func (r *setting) Default() (*config.Config, error) {
	dst := &config.Config{}
	var err error
	defaultIO := io.DefaultIo()
	dst, err = defaultIO.GetConfig()
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func NewSetting() domainSetting.Setting {
	return &setting{}
}
