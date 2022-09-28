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

func (m *Mac) Update(macAddress string) error {
	old, _ := m.Query()
	return dao.NewMac().Set(macAddress, old)
}

func (*Mac) Query() (string, error) {
	return dao.NewMac().Query()
}
