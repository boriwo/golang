//
// Golang Workshop 2024
//

package foo

import "github.com/rs/zerolog/log"

var (
	ExportedVar int = 42
	privateVar  int = 76
)

func ExportedFunction() string {
	log.Print("running exported function")
	return "hello world!!!"
}

func privateFunction() string {
	return "very private!!!"
}
