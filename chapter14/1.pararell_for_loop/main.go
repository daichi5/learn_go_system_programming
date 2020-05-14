package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"task1",
		"task2",
		"task3",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task string) {
			fmt.Println(task)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
