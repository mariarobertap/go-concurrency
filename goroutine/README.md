<h1>Goroutine</h1>

Se você já utilizou threads em outra linguagem, pode supor que as goroutines são similares as threads,
uma das diferença é que as goroutines são muito mais leves.
Uma goroutine nada mais é do que uma atividade que é execudada de forma concorrente. 
Você pode ter diversas atividades em um mesmo programa rodando de forma concorrente.

Exemplo 
```go
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
```
