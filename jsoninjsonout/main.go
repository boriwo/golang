//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Foo string `json:"foo,omitempty"`
}

func main() {

	str :=
		`{
			"Foo" : "Bar" 
		 }`

	// parse json string

	var m map[string]string

	err := json.Unmarshal([]byte(str), &m)

	if err != nil {
		fmt.Println(err)
	}

	// now you can reach inside the object

	if m["Foo"] != "Bar" {
		fmt.Println("unexpected data")
	}

	// serialize json to get the original string back

	buf, err := json.MarshalIndent(m, "", "\t")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", string(buf))
	}

	// alternatively, you can initialize a struct (or a map) in a single line

	m = map[string]string{"Foo": "Bar"}

	// serialization hints in structs

	str2 :=
		`{
			"foo" : "Bar" 
		 }`

	var msg Message

	err = json.Unmarshal([]byte(str2), &msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg.Foo)
	}
}
