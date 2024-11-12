//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"sync"
)

type StringMap interface {
	Get(string) string
	Set(string, string)
}

type SafeStringMap struct {
	sync.Mutex
	m map[string]string
}

func NewSafeStringMap(size int) StringMap {
	return &SafeStringMap{
		m: make(map[string]string, size),
	}
}

func (s *SafeStringMap) Get(key string) string {
	s.Lock()
	defer s.Unlock()
	return s.m[key]
}

func (s *SafeStringMap) Set(key string, value string) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func main() {
	m := NewSafeStringMap(100)
	var wg sync.WaitGroup
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			m.Set("foo", "bar")
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			m.Get("foo")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
