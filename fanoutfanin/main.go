//
// Golang Workshop 2024
//

// choose number of jobs and number of workers
// fan-out jobs, fan-in results
// synchronize workers with channels and wait groups
// ensure workers are closed at the end by closing channels or using a done channel

// Fanout refers to the pattern where a single goroutine (the sender) sends messages to
// multiple goroutines (the receivers) simultaneously. This pattern is commonly used when
// you have a single producer that generates data, and you want to distribute this data to
// multiple consumers for processing concurrently.

// Fanin is the opposite pattern, where multiple goroutines send messages to a single goroutine,
// which collects and processes these messages. This pattern is useful when you have multiple
// producers, and you want to consolidate their output into a single channel for further
// processing.

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	sum := 0

	for a := 1; a <= numJobs; a++ {
		sum += <-results
	}

	fmt.Println("sum of all jobs is ", sum)
}
