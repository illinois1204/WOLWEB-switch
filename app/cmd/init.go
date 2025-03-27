package cmd

import (
	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/repository"
	"github.com/illinois1204/WOLWEB-switch/app/service"
)

func RunAppInitialization() {
	constants.AppEnv.Load()
	fileListSet, _, lastWrittenFile := service.ListStoreFiles()
	var lastIndex uint = 0
	if lastWrittenFile != "" {
		var err error
		lastIndex, err = service.ExtractFileNameIndex(lastWrittenFile)
		if err != nil {
			panic(err)
		}
	}
	service.MakeCounter(int64(lastIndex))
	repository.DeviceStorage.Load(fileListSet)
}
