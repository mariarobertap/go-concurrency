package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {

	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {

	var mutex sync.Mutex

	wg.Add(2)
	updateMessage("Message 2", &mutex)

	updateMessage("Message 1", &mutex)

	wg.Wait()

	fmt.Println(msg)

}
