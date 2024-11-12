//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"github.com/sebdah/goldie/v2"
	"testing"
)

func TestMyFunction(t *testing.T) {
	g := goldie.New(t)
	myOutput := MyFunction()
	buf, _ := json.MarshalIndent(myOutput, "", "\t")
	g.Assert(t, "my_function_output", buf)
}
