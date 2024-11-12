//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	done := make(chan bool, 0)
	work := make(chan string, 0)

	go func() {
		for {
			fmt.Println("waiting for work")
			select {
			case w := <-work:
				// do work
				fmt.Println("doing work: ", w)
			case <-done:
				// cleanup and exit go routine in a clean way
				fmt.Println("cleaning up")
				return
			}
		}
	}()

	// create a channel to receive os.Signal values

	sigs := make(chan os.Signal, 1)

	// notify sigs when a SIGINT (Ctrl-C) is received

	signal.Notify(sigs, syscall.SIGINT)

	// wait for signal to come in

	<-sigs

	// now cleanup

	done <- true
	time.Sleep(3 * time.Second)
	fmt.Println("exiting")
}
