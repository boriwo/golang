//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

func doWork() error {
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	// deferred statement will be executed when surrounding function returns
	// useful for releasing resource when function has many return paths
	defer file.Close()
	data := map[string]string{"Foo": "Bar", "Hello": "World"}
	buf, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	num, err := file.Write(buf)
	if err != nil {
		return err
	}
	if num < len(buf) {
		return err
	}
	_, err = file.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}

func do(m map[string]string, lock sync.RWMutex) error {
	lock.Lock()
	defer lock.Unlock()
	val, ok := m["foo"]
	if !ok {
		return errors.New("missing key")
	}
	m["foo"] = val + "abc"
	return nil
}

func main() {
	err := doWork()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")
}
