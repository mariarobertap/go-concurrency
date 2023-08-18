package main

import "fmt"

func main() {

	messagesChannel := make(chan string)//bidirectional channel

	go sendMessage(messagesChannel,"Oi do sendMessage")//send channel
	receiveMessage(messagesChannel)//receive channel
}

func sendMessage(cs chan<- string,message string){
	cs<-message
}

func receiveMessage(cr <-chan string){
	fmt.Printf(<-cr)
}
