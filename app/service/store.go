package service

import (
	"os"
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
		panic(err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	var first, last string
	if len(files) > 0 {
		first = files[0]
		last = files[len(files)-1]
	}

	return files, first, last
}

func GetLastFileIndex(filename string) int {
	index := strings.Split(filename, ".json")[0]
	numIndex, err := strconv.Atoi(index)
	if err != nil {
		panic(err)
	}
	return numIndex
}
