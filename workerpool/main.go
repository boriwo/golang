//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	NUM_JOBS = 10
)

type Job struct {
	Task string
}

func NewJob(Task string) *Job {
	return &Job{Task}
}

type Worker struct {
	JobQueue    chan *Job
	DoneChannel chan bool
	Id          string
}

func NewWorker(JobQueue chan *Job, DoneChannel chan bool, Id string) *Worker {
	return &Worker{
		JobQueue:    JobQueue,
		DoneChannel: DoneChannel,
		Id:          Id,
	}
}

func (w *Worker) Run() {
	go func() {
		fmt.Printf("worker %s started\n", w.Id)
		for {
			select {
			case j := <-w.JobQueue:
				w.DoJob(j)
			case <-w.DoneChannel:
				fmt.Printf("worker %s finished\n", w.Id)
				return
			case <-time.After(1 * time.Second):
				fmt.Printf("no new job within 1 second\n")
			}
		}
	}()
}

func (w *Worker) Shutdown() {
	w.DoneChannel <- true
}

func (w *Worker) DoJob(j *Job) {
	if j == nil {
		return
	}
	fmt.Printf("running job %s\n", j.Task)
	time.Sleep(1 * time.Second)
}

func main() {
	jobs := make(chan *Job, NUM_JOBS)
	pool := make([]*Worker, 0)
	for i := 0; i < NUM_JOBS; i++ {
		pool = append(pool, NewWorker(jobs, make(chan bool), strconv.Itoa(i)))
	}
	for i := 0; i < NUM_JOBS; i++ {
		pool[i].Run()
	}
	for i := 0; i < NUM_JOBS; i++ {
		jobs <- NewJob(strconv.Itoa(i))
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < NUM_JOBS; i++ {
		pool[i].Shutdown()
	}
	time.Sleep(3 * time.Second)
}
