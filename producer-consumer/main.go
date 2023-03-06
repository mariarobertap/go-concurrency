package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaId int) *PizzaOrder {
	//Not using mutex because its only one go-routine trying to access the pizza data.
	pizzaId++

	if pizzaId <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order # %d\n", pizzaId)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++

		} else {
			pizzasMade++

		}
		total++

		fmt.Printf("Makind pizza #%d. It will take %d seconds....\n", pizzaId, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We ran out of ingredients for pizza #%d\n", pizzaId)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("***The cook quit while making the pizza #%d\n", pizzaId)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!\n", pizzaId)

		}

		return &PizzaOrder{
			pizzaNumber: pizzaId,
			message:     msg,
			success:     success,
		}
	}

	return &PizzaOrder{
		pizzaNumber: pizzaId,
	}
}

func pizzaria(pizzaJob *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaJob.data <- *currentPizza:

			case quitChan := <-pizzaJob.quit:
				close(pizzaJob.data)
				close(quitChan)
				return
			}
		}

	}

}

func (p *Producer) Close() error {
	ch := make(chan error)

	p.quit <- ch

	return <-ch
}

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("The Pizzaria is open for bussiness")
	color.Cyan("----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//Not using mutex because its only one go-routine trying to access the pizza data.
	go pizzaria(pizzaJob)

    //CONSUMER - consumes the pizzaria data
	//Similiar to a wait group. This for talks to the go routine to see the status of the pizzas
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas!")
			color.Cyan("----------------------------------")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel!")

			}
		}
	}

}
