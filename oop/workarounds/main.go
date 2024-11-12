//
// Golang Workshop 2024
//

package main

import "fmt"

type HashFct func() string

type Metric struct {
	Hash HashFct
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
	p.Metric.Hash = p.Hash
	p.Do()
}
