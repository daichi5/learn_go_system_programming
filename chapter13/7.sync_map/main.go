package main

import (
	"fmt"
	"sync"
)

func main() {
	swap := &sync.Map{}
	swap.Store("hello", "world")
	swap.Store(1, 2)

	swap.Delete("hello")
	swap.LoadOrStore(1, 4)
	value, ok := swap.Load("hello")
	fmt.Printf("key-%v value=%v exists?=%v\n", "hello", value, ok)
	value, ok = swap.Load(1)
	fmt.Printf("key-%v value=%v exists?=%v\n", 1, value, ok)
}
