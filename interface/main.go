//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func workOnStringer(item interface{ ToString() string }) {

}

func workOnAnything(item interface{}) {

	// determine type of variable using reflection

	fmt.Println("type of item is ", reflect.TypeOf(item))

	// type check with switch statement

	switch item.(type) {
	case map[string]string:

		// then safely perform a type cast

		m := item.(map[string]string)
		foo := m["Foo"]
		fmt.Printf("item is a string map and foo is %s\n", foo)

	case []string:
		fmt.Println("item is a string slice")
	default:
		fmt.Println("item is something else")
	}

	// type check by attempting type cast with comma ok idiom

	if _, ok := item.(string); ok {
		fmt.Println("item is a string")
	} else {
		fmt.Println("item is not a string")
	}

	// some functions happily operate on interfaces

	buf, err := json.MarshalIndent(item, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	// type conversion with type()

	str := string(buf)

	fmt.Println(str)

	// type cast to reach inside the interface

	foo := item.(map[string]string)["Foo"]

	fmt.Println(foo)
}

func main() {
	workOnAnything(map[string]string{"Foo": "Bar"})
}
