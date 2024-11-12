//
// Golang Workshop 2024
//

package main

func MyFunction() interface{} {
	m := map[string]interface{}{
		"Foo": "Bar",
		"Key": map[string]interface{}{
			"Hello": "World",
		},
	}
	return m
}
