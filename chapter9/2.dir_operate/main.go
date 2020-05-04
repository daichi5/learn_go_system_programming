package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("setting", 0755)
	printErr(err)
	err = os.MkdirAll("myapp/networksettings", 0755)
	printErr(err)

	defer func() {
		err = os.Remove("setting")
		printErr(err)
		err = os.RemoveAll("myapp")
		printErr(err)
	}()

	file, err := os.Create("textfile.txt")
	printErr(err)
	defer file.Close()
	io.WriteString(file, (strings.Repeat("a", 100)))
	file.Truncate(5)

	err = os.Rename("textfile.txt", "newfile.txt")
	printErr(err)
	err = os.Rename("./newfile.txt", "myapp/newfile.txt")
	printErr(err)
	fmt.Println("operated")
}
