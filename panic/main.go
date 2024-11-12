//
// Golang Workshop 2024
//

package main

import (
	"fmt"
)

func doDivByZero() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	div := 0
	inf := 10 / div
	fmt.Printf("infinity is %d\n", inf)
}

func doNilPointer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	var m map[string]string
	m["foo"] = "bar"
}

func doExplicitPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	panic("explicit panic")
}

func main() {
	doDivByZero()
	doNilPointer()
	doExplicitPanic()
}
