//
// Golang Workshop 2024
//

package main

import (
	"fmt"
	"strconv"
)

type Shape interface {
	Draw() error
}

type Color struct {
	R, G, B int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Color
	Point
	Radius int
}

func (c *Circle) Draw() error {
	// draw circle...
	fmt.Println("circle with radius " + strconv.Itoa(c.Radius) + " at point " + strconv.Itoa(c.X) + ", " + strconv.Itoa(c.Y))
	return nil
}

func main() {
	var circle Shape
	circle = &Circle{Color{10, 10, 10}, Point{200, 300}, 20}
	circle.Draw()
}
