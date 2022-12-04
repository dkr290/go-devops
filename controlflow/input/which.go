package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Which() {
	args := returnArgs()

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	var found bool = false
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		//does it exists
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			found = true
			// is it regular file
			if mode.IsRegular() {
				// is it executable
				if mode&0111 != 0 {
					fmt.Println(fullPath)

					continue
				}
			}
		}

	}
	if !found {
		fmt.Println("Could not find executable", file, "in PATH path")
		return
	}
}
