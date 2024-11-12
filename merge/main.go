//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"sync"
)

func merge(channels ...chan int) chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	c1 := make(chan int, 3)
	c2 := make(chan int, 3)

	merged := merge(c1, c2)

	go func() {
		for _, x := range []int{1, 2, 3} {
			c1 <- x
		}
		close(c1)
	}()

	go func() {
		for _, x := range []int{4, 5, 6} {
			c2 <- x
		}
		close(c2)
	}()

	for n := range merged {
		fmt.Println(n)
	}
}
