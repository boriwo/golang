//
// Golang Workshop 2024
//

package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	val := ctx.Value("foo")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("task completed successfully ", val)
	case <-ctx.Done():
		fmt.Println("task was canceled", val)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx = context.WithValue(ctx, "foo", "bar")
	defer cancel()
	fmt.Println("starting task")
	go longRunningTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
