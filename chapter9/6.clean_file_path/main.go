package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	absPath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(absPath)

	relpath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relpath)

	fmt.Println(filepath.Match("image-*.png", "image-100.png"))

	files, err := filepath.Glob("./*.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
