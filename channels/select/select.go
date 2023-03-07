package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----- SELECT WITH CHANNELS -----")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	for {
		//Note that case 1 and 2, and 3 and 4 are exactly the same. if this happens, go will choose randomly.
		select {
		case s1 := <-ch1:
			fmt.Println("CASE 1 - ", s1)
		case s2 := <-ch1:
			fmt.Println("CASE 2 - ", s2)
		case s3 := <-ch2:
			fmt.Println("CASE 3 - ", s3)
		case s4 := <-ch2:
			fmt.Println("CASE 4 - ", s4)
			//default - default case can also be used to avoid deadlocks
		}
	}

}

func server1(ch chan string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "SERVER1. This is a message sent from server 1."
	}
}
func server2(ch chan string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "SERVER 2. This is a message sent from server 2."
	}
}
