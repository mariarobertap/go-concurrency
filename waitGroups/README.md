# WaitGroups

Levando em conta que GoRoutines são funções que são executadas em paralelo como uma thread, seu programa pode finalizar antes que sua GoRoutine finalize.
Para contornar isso, existe os waitGroups, utilizados para informar ao seu programa que ele deve aguardar por uma ou mais GoRoutines finalizar


  - WaitGroup espera por uma quantidade de goroutines finalizar. Impedindo que seu programa finalize antes que as go routines terminem de rodar. 
  - As go routines são responsáveis por avisar que ela está finalizada.

# Exemplo:
```go
var wg sync.WaitGroup

func hello() {
	//Goroutine Avisando ao WaitGroup que finalizou.
	defer wg.Done()
	fmt.Println("hello from a goroutine!")
}

func main() {
	//wait group deve esperar por 1 goroutine finalizar
	wg.Add(1)
	//Go routine
	go hello()
	//Esperando Goroutine finalizar.
	wg.Wait()
}
```



