//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func Expensive(num int) {
	fmt.Printf("%d\n", num)
	time.Sleep(1 * time.Second)
}

func SequentialExecution() {
	for i := 0; i < 10; i++ {
		Expensive(i)
	}
}

func ConcurrentExecution() {
	for i := 0; i < 10; i++ {
		// go keyword runs function as light-weight thread
		go Expensive(i)
	}
}

func ConcurrentExecutionAnonymousFlawed() {
	for i := 0; i < 10; i++ {
		// closure over variable i causes problems due to concurrent execution
		go func() {
			fmt.Printf("%d\n", i)
			time.Sleep(1 * time.Second)
		}()
	}
	//time.Sleep(2 * time.Second)
}

func ConcurrentExecutionAnonymous() {
	for i := 0; i < 10; i++ {
		// better to pass data into function as parameter
		go func(num int) {
			fmt.Printf("%d\n", num)
			time.Sleep(1 * time.Second)
		}(i)
	}
	//time.Sleep(2 * time.Second)
}

func ConcurrentExecutionAnonymousWithWaitGroup() {
	// if we don't wait, the function will return before all the work is done
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Printf("%d\n", num)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	ConcurrentExecution()
	time.Sleep(3 * time.Second)
}
