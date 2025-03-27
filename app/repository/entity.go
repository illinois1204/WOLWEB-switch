package repository

type Device struct {
	Name string `form:"name" json:"name"`
	Mac  string `form:"mac" json:"mac"`
	Port uint16 `form:"port" json:"port"`
}

type IDeviceSet struct {
	Id uint
	Device
}

type DeviceLoadStub map[uint]Device

var DeviceStorage = make(DeviceLoadStub)
