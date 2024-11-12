//
// Golang Workshop 2024
//

package main

import (
	"fmt"

	"github.com/boriwo/golang/packages/foo"
)

func main() {
	fmt.Println(foo.ExportedFunction())
	fmt.Println(foo.ExportedVar)
}