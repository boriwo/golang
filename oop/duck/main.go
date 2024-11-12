//
// Golang Workshop 2024
//

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Chicken struct {
}

func (c *Chicken) Speak() string {
	return "i am a chicken"
}

type Duck struct {
}

func (d *Duck) Speak() string {
	return "i am a duck"
}

func main() {
	animals := []Speaker{new(Chicken), new(Duck)}
	for _, a := range animals {
		fmt.Printf("%s\n", a.Speak())
	}
}
