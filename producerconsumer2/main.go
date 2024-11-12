//
// Golang Workshop 2024
//

package main

import (
	"fmt"
)

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		fmt.Printf("produced %d\n", i)
		ch <- i
	}
}

func consumer(ch chan int, exit chan struct{}) {
	for num := range ch {
		fmt.Printf("consumed %d\n", num)
	}
	close(exit)
}

func main() {
	ch := make(chan int)
	exit := make(chan struct{})
	go producer(ch)
	go consumer(ch, exit)
	<-exit
}
