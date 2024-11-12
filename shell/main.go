//
// Golang Workshop 2024
//

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("%s", stdout.String())
	fmt.Printf("%s", stderr.String())
}
