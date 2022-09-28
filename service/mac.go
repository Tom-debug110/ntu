package service

import (
	"ntu/dao"
	"sync"
)

type Mac struct{}

var (
	macServiceInstance *Mac
	macServiceOnce     sync.Once
)

func NewMac() *Mac {
	macServiceOnce.Do(func() {
		macServiceInstance = &Mac{}
	})

	return macServiceInstance
}

func (*Mac) Update(macAddress string) error {
	return dao.NewMac().Set(macAddress)
}

func (*Mac) Query() (string, error) {
	return dao.NewMac().Query()
}
