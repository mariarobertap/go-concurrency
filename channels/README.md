<h1>O que são Channels?</h1> 

![Gravação de Tela 2023-03-08 às 10 15 26](https://user-images.githubusercontent.com/75685022/224519931-89e5e1c8-ede3-400e-8d08-0038bf317029.gif)



Um channel é um sistema de comunicação que permite uma go-routine enviar valores a outra go-routine. <br>
O channel nada mais é do que um cano utilizado para transmitir informações entre go-routines.

<h3>Como criar um channel?</h3> 

```go
ch := make(chan int)
```



"Don't communicate by sharing memory; share memory by communicating. (R. Pike)"
