package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// StringFuture is..
type StringFuture struct {
	receiver chan string
	cache    string
}

// NewStringFuture is...
func NewStringFuture() (*StringFuture, chan string) {
	f := &StringFuture{
		receiver: make(chan string),
	}
	return f, f.receiver
}

// Get is...
func (f *StringFuture) Get() string {
	r, ok := <-f.receiver
	if ok {
		close(f.receiver)
		f.cache = r
	}
	return f.cache
}

// Close is...
func (f *StringFuture) Close() {
	close(f.receiver)
}

func readFile(path string) *StringFuture {
	promise, future := NewStringFuture()
	go func() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			promise.Close()
		} else {
			future <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource *StringFuture) chan []string {
	promise := make(chan []string)
	go func() {
		var result []string
		for _, line := range strings.Split(futureSource.Get(), "\n") {
			if strings.HasPrefix(line, "func") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func countLines(futureSource *StringFuture) chan int {
	promise := make(chan int)
	go func() {
		promise <- len(strings.Split(futureSource.Get(), "\n"))
	}()
	return promise
}

func main() {
	futureSource := readFile("hoge.txt")
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
	fmt.Println(<-countLines(futureSource))
}
