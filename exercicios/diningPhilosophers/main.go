package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosofers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

//
var hunger = 1 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second
var listPhilosofers []string

func main() {
	fmt.Println("Dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	dine()

	fmt.Println("The table is empty.")

}

func dine() {

	wg := &sync.WaitGroup{}
	wg.Add(len(philosofers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosofers))
	list := &sync.Mutex{}

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosofers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//Start the meal.
	for i := 0; i < len(philosofers); i++ {
		//Fire of a goroutine for the current philosopher
		go diningProblem(philosofers[i], wg, forks, seated, list)

	}

	wg.Wait()
	fmt.Println(listPhilosofers)

}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup, list *sync.Mutex) {
	defer wg.Done()

	//seat the philoshphers at the table

	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	seated.Done()
	//Wait until everyone is seated.
	seated.Wait()

	for i := 0; i < hunger; i++ {
		//get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("    %s Takes the rifht fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("    %s Takes the left fork.\n", philosopher.name)

		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("    %s Takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("    %s Takes the rifht fork.\n", philosopher.name)
		}

		fmt.Printf("    %s has both forks, is eating..\n", philosopher.name)
		time.Sleep(eatTime)
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("   %s Is done eating.\n", philosopher.name)

	}
	list.Lock()
	listPhilosofers = append(listPhilosofers, philosopher.name)
	list.Unlock()

}
