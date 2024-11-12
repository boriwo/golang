//
// Golang Workshop 2024
//

package foo

var (
	ExportedVar int = 42
	privateVar  int = 76
)

func ExportedFunction() string {
	return "hello world!!!"
}

func privateFunction() string {
	return "very private!!!"
}
