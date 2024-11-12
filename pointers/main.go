//
// Golang Workshop 2024
//

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Node struct {
	Value int
	Next  *Node
}

func InsertNode(node **Node, value int) {
	newNode := &Node{Value: value}
	if *node != nil {
		newNode.Next = (*node).Next
		(*node).Next = newNode
	} else {
		*node = newNode
	}
}

func pointerBasics() {

	var p *int
	var x int = 10

	p = &x

	fmt.Println(*p)

	*p = 42

	fmt.Println(*p)
}

func setName(person *Person, name string) {
	person.Name = name
}

func (person *Person) SetName(name string) {
	person.Name = name
}

func main() {
	pointerBasics()
	p := new(Person)
	setName(p, "Boris")
	fmt.Println("name is ", p.Name)
	p.SetName("Michael")
	fmt.Println("name is ", p.Name)
}
