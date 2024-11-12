//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"sync"
)

// maps and slices

func main() {

	// allocate map with make()

	m1 := make(map[string]int)

	// ...or initialize map in one line

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// read and write elements

	m1["one"] = 1
	val := m2["too"]
	fmt.Printf("map element two is %d\n", val)

	// iterate over all elements

	for k, v := range m2 {
		fmt.Printf("map element %s is %d\n", k, v)
	}

	// comma ok idiom to check for presence of values

	val, ok := m2["two"]

	if ok {
		fmt.Printf("map element is present, value is %d\n", val)
	} else {
		fmt.Printf("map element is not present, value is %d\n", val)
	}

	// remove element with delete()

	delete(m1, "one")

	// check length with len

	if len(m2) > 2 {
		fmt.Println("m2 has more than 2 elements")
	}

	// maps are not thread-safe, protect with lock for concurrent access

	lock := sync.RWMutex{}

	lock.Lock()
	m1["one"] = 1
	lock.Unlock()

	// slices are arrays that can dynamically change size

	slice := []int{4, 5, 6, 7}

	slice2 := make([]int, 4)

	slice[0] = 3

	slice2[0] = 3

	// add elements with append(), must reassign reference because it may change if memory has to be reallocated

	slice = append(slice, 8)

	// iterate of slice

	for idx, val := range slice {
		fmt.Printf("element %d has value %d\n", idx, val)
	}

	// use len to determine number of elements

	if len(slice) > 3 {
		fmt.Println("slice has more than 3 elements")
	}

	// take a slice of a slice

	slice = slice[2:4]

	for idx, val := range slice {
		fmt.Printf("element %d has value %d\n", idx, val)
	}

	// remove element

	slice = []int{1, 2, 3, 4, 5}
	index := 2
	slice = append(slice[:index], slice[index+1:]...)

	for idx, val := range slice {
		fmt.Printf("element %d has value %d\n", idx, val)
	}
}
