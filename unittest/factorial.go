//
// Golang Workshop 2024
//

package main

import (
	"sync"
)

func FactorialA(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FactorialA(n-1)
}

func FactorialB(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialC(cache map[int]int, n int) int {
	if result, ok := cache[n]; ok {
		return result
	}
	result := FactorialB(n)
	cache[n] = result
	return result
}

func RaceCondition(n int) int {
	var wg sync.WaitGroup
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i * 2
		}(i)
	}
	wg.Wait()
	return m[n]
}

func main() {
}
