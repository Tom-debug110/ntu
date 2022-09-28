package dao

import (
	"ntu/model"
	"sync"
)

type Mac struct{}

var (
	MacDaoInstance *Mac
	MacDaoOnce     sync.Once
)

func NewMac() *Mac {
	MacDaoOnce.Do(func() {
		MacDaoInstance = &Mac{}
	})
	return MacDaoInstance
}

func (*Mac) Set(macAddress string) error {
	return db.Model(model.MacAddress{}).Update("id", macAddress).Error
}

func (*Mac) Query() (string, error) {
	var m model.MacAddress
	err := db.Model(model.MacAddress{}).First(&m).Error
	return m.ID, err
}
