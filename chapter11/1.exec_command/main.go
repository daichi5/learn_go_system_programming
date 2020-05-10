package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("  Pid: %d\n", state.Pid())
	fmt.Printf("  System: %d\n", state.SystemTime())
	fmt.Printf("  User: %d\n", state.UserTime())
	io.Copy(os.Stdout, bytes.NewBuffer(output))
}
