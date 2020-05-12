package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Accept Ctrl + C for 5 second")
	time.Sleep(time.Second * 5)

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	fmt.Println("Ignore Ctrl + C for 10 second")
	time.Sleep(time.Second * 100)
}
