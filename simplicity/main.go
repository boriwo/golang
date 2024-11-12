//
// Golang Workshop 2024
//

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	// variable declaration and initialization in one line

	var message string = "hello"

	// or even shorter

	msg := "world"

	// variable declaration and initialization in one line

	var a int = 42

	// or even shorter

	b := 84

	print(a, b)

	// no semicolons

	// string concatenation

	fmt.Println(message + " " + msg)

	// variadic functions

	fmt.Println(message, " ", msg)

	// no parentheses in control structures

	if 1 == 1 {
		fmt.Println("one is still one")
	}

	// implicit variable definitions and C-inspired syntax

	for i := 1; i <= 3; i++ {
		fmt.Printf("i have seen this %d times\n", i)
	}

	// infinite loops and break

	i := 1
	for {
		if i > 3 {
			break
		}
		fmt.Printf("again, i have seen this %d times\n", i)
		i++
	}

	// switch without break

	v := "foo"

	switch v {
	case "foo":
		fmt.Println("v is foo")
	case "bar":
		fmt.Println("v is bar")
	default:
		fmt.Println("v is none of the above")
	}

	// functions can have more than one return value, useful for error handling and many other things

	num := -64.0
	root, err := Sqrt(num)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("square root of %f is %f\n", num, root)
	}

	// ...or if you are brave, ignore the error

	root, _ = Sqrt(num)

	// instantiate a struct in one line

	foo := NewFoo("abc", 123)

	foo = &Foo{
		"abc",
		123,
	}

	// call member function on a struct

	foo.num = 23
	fmt.Println(foo.num)

	foo = new(Foo)
	foo.num = 123
	foo.str = "abc"

	foo.ToString()

	foo.SetNum(42)

	do()
}

func NewFoo(str string, num int) *Foo {
	return &Foo{
		str,
		num,
	}
}

// defining a function with multiple return values

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("negative numbers not allowed")
	}
	return math.Sqrt(f), nil
}

// defining a struct

type Foo struct {
	str string
	num int
}

// defining functions on structs

func (f *Foo) ToString() {
	fmt.Printf("%s %d\n", f.str, f.num)
}

// defining functions on structs

func (f *Foo) SetNum(num int) {
	f.num = num
}

func (f *Foo) GetNum() int {
	return f.num
}

/*func increment(num *int) {
	*num = *num + 1
}

func do() {
	var i int = 10
	increment(&i)
	fmt.Println(i)
}*/

func increment(num int) int {
	return num + 1
}

func do() {
	var i int = 10
	i = increment(i)
	fmt.Println(i)
}

type Inner struct {
	Name string
}

type Outer struct {
	Inner
	Name string
}

func innerouter() {

	o := Outer{
		Inner: Inner{Name: "Inner Name"},
		Name:  "Outer Name",
	}

	fmt.Println(o.Name)
	fmt.Println(o.Inner.Name)
}
