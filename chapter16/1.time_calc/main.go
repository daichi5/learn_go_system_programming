package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Add(3 * time.Hour))
	fmt.Println(time.Now().Round(time.Hour))
	fmt.Println(time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local))
	fmt.Println(time.Date(2017, time.August, 26, 11, 50, 30, 0, time.UTC))
}
