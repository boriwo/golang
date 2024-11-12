//
// Golang Workshop 2024
//

package main

import "fmt"

func mapFunc(input []int, mapper func(int) int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func filterFunc(input []int, filter func(int) bool) []int {
	result := []int{}
	for _, v := range input {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	doubled := mapFunc(numbers, func(x int) int {
		return x * 2
	})

	mf := func(x int) int {
		return x * 2
	}

	doubled = mapFunc(numbers, mf)

	evens := filterFunc(numbers, func(x int) bool {
		return x%2 == 0
	})

	fmt.Println("doubled: ", doubled)
	fmt.Println("even: ", evens)
}
