package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var id int
var ids = []int{}

func generateID(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	result := id
	ids = append(ids, result)
	return result
}

func main() {
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateID(&mutex))
		}()
	}
	time.Sleep(time.Second)
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	fmt.Println(ids)
}
