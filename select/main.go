//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"time"
)

func doJob(job string) {
}

func main() {

	jobQueue1 := make(chan string)
	jobQueue2 := make(chan string)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case j := <-jobQueue1:
				doJob(j)
			case j := <-jobQueue2:
				doJob(j)
			case <-done:
				fmt.Println("shutting down")
				return
			case <-time.After(60 * time.Second):
				fmt.Println("alert: no new job within 60 seconds")
			}
		}
	}()
}
