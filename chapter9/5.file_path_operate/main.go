package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	joinedPath := filepath.Join("foo", "hoge", "fuga")
	fmt.Printf("Joined path: %s\n", joinedPath)

	dir, name := filepath.Split(joinedPath)
	fmt.Printf("Splitted path, dir: %s, name: %s \n", dir, name)

	fragments := strings.Split(joinedPath, string(filepath.Separator))
	for _, fragment := range fragments {
		fmt.Println(fragment)
	}
	var pathEnv string = "/usr:/usr/local:/usr/bin"
	pathList := filepath.SplitList(pathEnv)
	for _, path := range pathList {
		fmt.Println(path)
	}
}
