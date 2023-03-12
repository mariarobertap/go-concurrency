package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	var balance sync.Mutex

	income := []Income{
		{
			Source: "Main job",
			Amount: 1000,
		},
		{
			Source: "Investments",
			Amount: 500,
		},
	}

	wg.Add(len(income))
	for i := 0; i < len(income); i++ {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 0; week < 4; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

			}

		}(i, income[i])

	}
	wg.Wait()

	fmt.Printf("Total amount earned $ %d", bankBalance)

}
