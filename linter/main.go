//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := map[string]string{"Hello": "Golang"}

	if m != nil && len(m) > 0 {
		m["Hello"] = "World"
	}

	buf, _ := json.Marshal(m)

	fmt.Printf("%d\n", string(buf))
}
