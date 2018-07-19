package main

import (
	"fmt"
	"os"
	"os/exec"
)

func myreader() []string {
	var r []string
	for i := range os.Args {
		if i > 2 {
			r = append(r, os.Args[i])
		}
	}
	return r
}

func myexe(s ...string) {
	app := s[0]
	args := s[1:len(s)]
	exec.Command(app, args...).Run()
}
func main() {
	if len(os.Args) < 4 {
		fmt.Println("compileIfChanged.exe program.go program.exe go build program.go")
		return
	}
	file1, _ := os.Lstat(os.Args[1])
	file2, _ := os.Lstat(os.Args[2])
	if file1.ModTime().After(file2.ModTime()) {
		c := []string{"cmd", "/c"}
		c2 := myreader()
		c = append(c, c2...)
		myexe(c...)
		fmt.Println(file1.Name() + " is modified after " + file2.Name())
	}
}
