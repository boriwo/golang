//
// Golang Workshop 2024
//

package main

type Conn struct {
}

type Result struct {
}

func (c *Conn) DoQuery(query string) Result {
	return Result{}
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result, len(conns))
	for _, conn := range conns {
		go func(c Conn) {
			ch <- c.DoQuery(query)
		}(conn)
	}
	return <-ch
}

func main() {
	Query(make([]Conn, 3), "select foo from bar")
}
