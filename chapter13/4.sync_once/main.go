package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("initialized")
}

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
