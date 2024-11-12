package main

import (
	"fmt"
)

func main() {

	const s = "你好世界"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	for idx, rune := range s {
		fmt.Printf("%c starts at %d\n", rune, idx)
	}
}
