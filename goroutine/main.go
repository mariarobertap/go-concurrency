package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printSomething("Hello from goroutine 1", &wg)
	fmt.Println("Hello from main-goroutine.")

	wg.Wait()

}
