//
// Golang Workshop 2024
//

package main

import "fmt"

type Metric struct {
}

func (m *Metric) Hash() string {
	return "_"
}

func (m *Metric) Do() {
	fmt.Println(m.Hash())
}

type Plugin struct {
	Metric
}

func (p *Plugin) Hash() string {
	return "123"
}

func main() {
	p := new(Plugin)
	p.Do()
}
