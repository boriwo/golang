//
// Golang Workshop 2024
//

package main

import "fmt"

func SumInts(a []int64) int64 {
	var s int64
	for _, v := range a {
		s += v
	}
	return s
}

func SumFloats(a []float64) float64 {
	var s float64
	for _, v := range a {
		s += v
	}
	return s
}

func Sum[V int64 | float64](a []V) V {
	var s V
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	a := []int64{1, 2, 3}
	s := Sum(a)
	fmt.Println(s)
}
