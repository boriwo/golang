//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Job struct {
	Task string
}

func NewJob(Task string) *Job {
	return &Job{Task}
}

type DoJob func(*Job) string

type Worker struct {
	jobQueue chan *Job
	done     chan struct{}
	id       string
	do       DoJob
	pool     *Pool
}

func NewWorker(id string, do DoJob, pool *Pool) *Worker {
	return &Worker{
		jobQueue: make(chan *Job),
		done:     make(chan struct{}),
		id:       id,
		do:       do,
		pool:     pool,
	}
}

func (w *Worker) Run() *Worker {
	go func() {
		for {
			select {
			case j := <-w.jobQueue:
				fmt.Println("at", time.Now().UnixMilli(), "worker", w.id, "task", j.Task, "started")
				w.do(j)
				fmt.Println("at", time.Now().UnixMilli(), "worker", w.id, "task", j.Task, "finished")
				w.pool.available <- w
			case <-w.done:
				fmt.Println("shutdown", w.id)
				return
			}
		}
	}()
	return w
}

type Pool struct {
	available chan *Worker
	jobsQueue chan *Job
}

func NewPool(size int, Do DoJob, JobQueue chan *Job) *Pool {
	p := &Pool{
		available: make(chan *Worker, size),
		jobsQueue: JobQueue,
	}
	for i := 0; i < size; i++ {
		p.available <- NewWorker(strconv.Itoa(i), Do, p).Run()
	}
	return p
}

func (p *Pool) Launch() {
	go func() {
		for j := range p.jobsQueue {
			w := <-p.available
			w.jobQueue <- j
		}
	}()
}

func (p *Pool) Shutdown() {
	go func() {
		for w := range p.available {
			w.done <- struct{}{}
		}
	}()
}

func main() {
	const NUM_JOBS = 30
	const NUM_WORKERS = 30
	jobs := make(chan *Job)
	do := func(j *Job) string {
		pauseSec := rand.Intn(5) + 1
		time.Sleep(time.Duration(pauseSec) * time.Second)
		return ""
	}
	pool := NewPool(NUM_WORKERS, do, jobs)
	pool.Launch()
	fmt.Println("sending work")
	for i := 0; i < NUM_JOBS; i++ {
		fmt.Println("at", time.Now().UnixMilli(), "sending job")
		jobs <- NewJob(strconv.Itoa(i))
	}
	fmt.Println("shutting down")
	pool.Shutdown()
	time.Sleep(5 * time.Second)
}
