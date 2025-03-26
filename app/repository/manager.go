package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sync"

	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/service"
)

func Write(object Device) {
	jsonBytes, err := json.MarshalIndent(object, "", constants.TabSpace)
	if err != nil {
		panic(err)
	}

	rowIndex := service.Next()
	fileName := fmt.Sprintf("%d.json", rowIndex)
	err = os.WriteFile((constants.StoreDir + "/" + fileName), jsonBytes, constants.FileWriteMode)
	if err != nil {
		panic(err)
	}

	DeviceStorage = append(DeviceStorage, DeviceLoadStub{
		Id:   uint(rowIndex),
		File: fileName,
		Data: object,
	})
}

func Load(files []string) {
	for _, file := range files {
		rawContent, err := os.ReadFile(constants.StoreDir + "/" + file)
		if err != nil {
			panic(err)
		}

		stub := Device{}
		if err := json.Unmarshal(rawContent, &stub); err != nil {
			panic(err)
		}

		rowIndex, err := service.ExtractFileNameIndex(file)
		if err != nil {
			panic(err)
		}

		DeviceStorage = append(DeviceStorage, DeviceLoadStub{
			Id:   rowIndex,
			File: file,
			Data: stub,
		})
	}
}

func ThreadLoad(files []string) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(_file string) {
			defer wg.Done()
			rawContent, err := os.ReadFile(constants.StoreDir + "/" + _file)
			if err != nil {
				// TODO: control this error and abort process (and next down checks)
				fmt.Println(err)
				return
			}

			stub := Device{}
			if err := json.Unmarshal(rawContent, &stub); err != nil {
				fmt.Println(err)
				return
			}

			rowIndex, err := service.ExtractFileNameIndex(file)
			if err != nil {
				fmt.Println(err)
				return
			}

			mu.Lock()
			DeviceStorage = append(DeviceStorage, DeviceLoadStub{
				Id:   rowIndex,
				File: file,
				Data: stub,
			})
			mu.Unlock()
		}(file)
	}
	wg.Wait()
	slices.SortFunc(DeviceStorage, func(a, b DeviceLoadStub) int { return int(a.Id) - int(b.Id) })
}
