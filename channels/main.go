package main

import (
	"fmt"
	"strings"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	var userInput string

	go shout(ping, pong)

	for {

		fmt.Print("->")

		fmt.Scanln(&userInput)

		if userInput == "q" {
			break
		}

		ping <- userInput

		response := <-pong

		fmt.Println(response)

	}
	close(ping)
	close(pong)
}

//Receiver and sender. receiver <-chan.  sender chan<-
func shout(receiver <-chan string, sender chan<- string) {

	for {
		s := <-receiver
		sender <- fmt.Sprintf("%s!!", strings.ToUpper(s))
	}
}
