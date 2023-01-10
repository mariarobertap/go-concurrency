package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {

	wg.Add(2)
	go updateMessage("Teste 1")
	go updateMessage("Teste 2")
	wg.Wait()

	fmt.Println(msg)
}
