//
// Golang Workshop 2024
//

package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func readOutput(reader io.Reader, ch chan string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
}

func main() {

	cmd := exec.Command("traceroute", "www.google.com")

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating stdout pipe: %s\n", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating stderr pipe: %s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %s\n", err)
		return
	}

	stdoutChan := make(chan string)
	stderrChan := make(chan string)

	go readOutput(stdoutPipe, stdoutChan)
	go readOutput(stderrPipe, stderrChan)

	go func() {
		for {
			select {
			case str, ok := <-stdoutChan:
				if ok {
					fmt.Println(str)
				}
			case str, ok := <-stderrChan:
				if ok {
					fmt.Println(str)
				}
			case <-time.After(time.Second * 60):
				fmt.Println("no output in 60 seconds, exiting...")
				return
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Printf("command finished with error: %v\n", err)
	}

	close(stdoutChan)
	close(stderrChan)
}
