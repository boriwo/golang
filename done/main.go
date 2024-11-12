//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("shutting down")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("working")
		}
	}
}

func main() {
	done := make(chan struct{})
	go worker(done)
	go worker(done)
	time.Sleep(3 * time.Second)
	close(done)
	//done <- struct{}{}
	//done <- struct{}{}
	time.Sleep(3 * time.Second)
}
