//
// Golang Workshop 2024
//

package main

import "fmt"

func sum(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}

/*func main() {
	s := sum(1, 2, 3)
	fmt.Printf("the some of 1, 2, 3 is %d\n", s)
}*/

func main() {
	s := func(num ...int) int {
		sum := 0
		for _, n := range num {
			sum += n
		}
		return sum
	}(1, 2, 3)
	fmt.Printf("the some of 1, 2, 3 is %d\n", s)
}

func main2() {
	f := func(num ...int) int {
		sum := 0
		for _, n := range num {
			sum += n
		}
		return sum
	}
	s := f(1, 2, 3)
	fmt.Printf("the some of 1, 2, 3 is %d\n", s)
}
