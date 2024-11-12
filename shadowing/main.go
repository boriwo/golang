//
// Golang Workshop 2024
//

package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)

	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}

	fmt.Println(x)
}
