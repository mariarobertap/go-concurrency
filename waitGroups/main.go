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
	words := []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
		"test6",
	}

	wg.Add(len(words))

	for i := 0; i < len(words); i++ {
		go printSomething(words[i], &wg)
	}

	wg.Wait()

}
