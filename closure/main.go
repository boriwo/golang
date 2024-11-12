//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"math"
)

type MathFunc func(x float64) float64

func Compose2(f, g MathFunc) MathFunc {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func Compose(f, g func(x float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func main() {
	print(Compose(math.Sin, math.Cos)(0.5))

	sum := 2
	func(x int) {
		sum += x
	}(3)
	fmt.Println(sum)
}
