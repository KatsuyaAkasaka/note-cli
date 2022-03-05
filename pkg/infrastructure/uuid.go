package infrastructure

import (
	"math/rand"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/uuid"
)

type uuidRepository struct{}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxMask = 0x3F // 63 0b111111
	uuidLen       = 6
)

func (r *uuidRepository) Gen() string {

	src := make([]byte, 1)
	buf := make([]byte, uuidLen)
	for i := 0; i < uuidLen; {
		if _, err := rand.Read(src); err != nil {
			panic(err)
		}
		idx := int(src[0] & letterIdxMask)
		if idx < len(letterBytes) {
			buf[i] = letterBytes[idx]
			i++
		}
	}
	return string(buf)
}

func NewUUIDRepository() uuid.Repository {
	return &uuidRepository{}
}
