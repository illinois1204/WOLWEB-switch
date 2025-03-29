package service

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/illinois1204/WOLWEB-switch/app/constants"
)

/*
Return 3 params:
 1. Array
 2. First element
 3. Last element
*/
func ListStoreFiles() ([]string, string, string) {
	entries, err := os.ReadDir(constants.StoreDir)
	if err != nil {
		entries = []os.DirEntry{}
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	sort.Slice(files, func(i, j int) bool {
		_Iindex, _ := ExtractFileNameIndex(files[i])
		_Jindex, _ := ExtractFileNameIndex(files[j])
		return _Iindex < _Jindex
	})

	var first, last string
	if len(files) > 0 {
		first = files[0]
		last = files[len(files)-1]
	}

	return files, first, last
}

func ExtractFileNameIndex(filename string) (uint, error) {
	index := strings.TrimSuffix(filename, ".json")
	numIndex, err := strconv.Atoi(index)
	if err != nil {
		return 0, err
	}
	return uint(numIndex), nil
}
