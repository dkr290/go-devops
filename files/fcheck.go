package main

import (
	"fmt"
	"os"
)

func checkFiles(f1, f2 string) {

	file, err := os.Stat(f1)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The file", f1, "does not exist")
			fmt.Println(f1)
		}
	}

	fmt.Println()
	file, err = os.Stat(f2)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The file", f2, "Does not exists")
		}
	}
	fmt.Printf("File name %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n", file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())

}
