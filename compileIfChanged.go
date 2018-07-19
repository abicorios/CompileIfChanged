package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Lstat(os.Args[1])
	file2, _ := os.Lstat(os.Args[2])
	if file1.ModTime().After(file2.ModTime()) {
		fmt.Println(file1.Name() + " is modified after " + file2.Name())
	}
}
