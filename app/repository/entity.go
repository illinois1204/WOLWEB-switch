package repository

type Device struct {
	Name string `form:"name" json:"name"`
	Mac  string `form:"mac" json:"mac"`
	Port uint16 `form:"port" json:"port"`
}

type DeviceLoadStub struct {
	Id   uint
	File string
	Data Device
}

var DeviceStorage []DeviceLoadStub
