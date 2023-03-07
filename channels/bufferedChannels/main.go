package main

import (
	"fmt"
	"time"
)

func main() {

	//This channel have a buffer of 10. I can put 10 values in this channel.
	//Everytime the listenToChan function get the value from the chan, i get one free space..
	ch := make(chan int, 10)

	go listenToChannel(ch)

	for i := 0; i < 100; i++ {

		fmt.Println("sendind data to channel", i)

		ch <- i
		fmt.Println("Sent data to channel", i)

	}

	close(ch)

}

func listenToChannel(ch chan int) {

	for {
		i := <-ch

		fmt.Println("Got", i, " from channel")

		time.Sleep(1 * time.Second)
	}
}
