<h1>O que são Channels?</h1> 

![Gravação de Tela 2023-03-08 às 10 15 26](https://user-images.githubusercontent.com/75685022/224519931-89e5e1c8-ede3-400e-8d08-0038bf317029.gif)



Um channel é um sistema de comunicação que permite uma go-routine enviar valores a outra go-routine. <br>
O channel nada mais é do que um cano utilizado para transmitir informações entre go-routines.

<h3>Como criar um channel?</h3> 

```go
ch := make(chan int)
```

<h3> Tipos de canais </h3>

Os canais podem ser especificos ou gerais ou seja eles podem ser unidirecionais ou bidirecionais.

<h4>Exemplo de canal bidirecional:</h4>

```go
ch := make(chan string)
```
Este canal pode ser convertido para um canal especifico(send ou receive)

<h4>Exemplo de canais unidirecionais:</h4>

```go
// canal de envio 
cs := ch<-
//canal de recebimento 
cr := <-ch
```
Canais unidirecionais não podem ser convertidos para bidirecional nem de um send para um receive depois de sua declaração

exemplos:
```go
//only bidirectional channels can become send or receive channels
package main
//it works :)
func main(){
	channel := make(chan string)
	channelSend :=  func() {
		 channel <- "Hello"
	}
	go channelSend() 
	channelReceive := <-channel
	fmt.Println(channelReceive)
}
	
```
```go
package main
//It doesn't work :(
func main(){
	cs := make(chan<-string)
	cr := make(<-chan string)

 	(<-chan string)(cs)
	(chan<- string)(cr)
	(chan string)(cs)
	(chan string)(cr)
}	
```

<h4>Exemplo usando funções: <h4>

```go 
package main

import "fmt"

func sendNumber(cs chan<- int64, number int64){
    cs <- number
}
func receiveAndReturnANumber(cr <-chan int64)int64{
    number := <-cr
    return number
}
func main(){
    ch := make(chan int64)
    go sendNumber(ch,1075897489377463545)
    number := receiveAndReturnANumber(ch)
    fmt.Println(number)
}
```

"Don't communicate by sharing memory; share memory by communicating. (R. Pike)"
