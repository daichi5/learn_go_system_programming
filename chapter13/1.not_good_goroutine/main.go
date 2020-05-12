package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
