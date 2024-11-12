//
// Golang Workshop 2024
//

package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("i can c you\n")
	C.puts(cs)
	C.free(unsafe.Pointer(cs))
}
