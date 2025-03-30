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

func Write(object Device) (uint, error) {
	jsonBytes, err := json.MarshalIndent(object, "", constants.TabSpace)
	if err != nil {
		return 0, err
	}

	rowIndex := service.Next()
	fileName := fmt.Sprintf("%d.json", rowIndex)
	err = os.WriteFile((constants.StoreDir + "/" + fileName), jsonBytes, constants.FileWriteMode)
	if err != nil {
		return 0, err
	}

	return uint(rowIndex), nil
}

func (s DeviceLoadStub) Add(key uint, data Device) {
	s[key] = data
}

func (s DeviceLoadStub) Load(files []string) {
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

		s[rowIndex] = stub
	}
}

func (s DeviceLoadStub) ThreadLoad(files []string) {
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
			s[rowIndex] = stub
			mu.Unlock()
		}(file)
	}
	wg.Wait()
}

func (s DeviceLoadStub) Update(id uint, payload Device) error {
	jsonBytes, err := json.MarshalIndent(payload, "", constants.TabSpace)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%d.json", id)
	err = os.WriteFile((constants.StoreDir + "/" + fileName), jsonBytes, constants.FileWriteMode)
	if err != nil {
		return err
	}

	s[id] = payload
	return nil
}

func (s DeviceLoadStub) Remove(id uint) {
	delete(s, id)
	os.Remove(fmt.Sprintf("%s/%d.json", constants.StoreDir, id))
}

func (s DeviceLoadStub) ToArray() []IDeviceSet {
	var arr []IDeviceSet
	for _key, _val := range s {
		arr = append(arr, IDeviceSet{Id: _key, Device: _val})
	}

	slices.SortFunc(arr, func(a, b IDeviceSet) int { return int(a.Id) - int(b.Id) })
	return arr
}
