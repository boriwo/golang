//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("get heap profile with: go tool pprof http://localhost:8080/debug/pprof/heap")
	go leakyFunction(wg)
	go goodFunction(wg)
	wg.Wait()
}

func goodFunction(wg sync.WaitGroup) {
	defer wg.Done()
	m := make(map[int]string)
	for i := 0; i < 10000000; i++ {
		m[i%100000] = "golang workshop"
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 0)
	for i := 0; i < 10000000; i++ {
		s = append(s, "golang workshop")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("running...")
		}
	}
}
