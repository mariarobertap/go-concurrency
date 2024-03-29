package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	//Esse é um buffered channel, com o tamanho == ao tamanho da capacidade de clientes da barbearia.
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day")

	shop.AddBarber("Frank")
	shopClosing := make(chan bool)
	closed := make(chan bool)

	//Verificando se 10s já se passaram para fechar a loja
	go func() {
		//Waiting until is time to close the store.
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true

	}()

	i := 0
	//Enviando clientes para a barbearia.
	go func() {
		for {
			//get a random number with average arrival date
			random := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(random)):
				shop.addClient(fmt.Sprintf("Maria #%d", i))
				i++

			}
		}
	}()

	//Esperando até que tudo estiver finalizado.
	<-closed

}

func (shop *BarberShop) AddBarber(barberName string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false

		color.Yellow(" %s Goes to the waiting room to check for clients.", barberName)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%s Takes a nap. No clients in the waiting room", barberName)
				isSleeping = true
			}

			//Verifica se o canal está fechado.
			//Se estiver, significa que a loja está fechada.
			//Mas como é um buffered channel, ainda pode ter dados no buffer.
			//Atendendo o critério do exercicio, que se a loja estiver fechada
			//ainda pode ter clientes na sala de espera.
			client, open := <-shop.ClientsChan

			if open {
				if isSleeping {
					color.Yellow("Client %s wakes up %s", client, barberName)
					isSleeping = false
				}

				shop.cutHair(barberName, client)

			} else {
				shop.sendBarberHome(barberName)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barberName, client string) {
	color.Green("%s is cutting %s hair", barberName, client)
	time.Sleep(shop.HairCutDuration)
}

func (shop *BarberShop) sendBarberHome(barberName string) {
	color.Cyan("%s is going home", barberName)
	shop.BarbersDoneChan <- true

}
func (shop *BarberShop) closeShopForDay() {

	color.Cyan("Closing shop for the day")
	//Fechando o canal para não aceitar novos clientes na sala de espera.
	//Os que ainda estáo esperando, continuam até serem atendidos.
	close(shop.ClientsChan)
	shop.Open = false

	for i := 0; i < shop.NumberOfBarbers; i++ {
		//esperando receber um valor aqui ==
		<-shop.BarbersDoneChan

	}

	close(shop.BarbersDoneChan)

	color.Green("-------------------")
	color.Green("Barber shop is closed")

}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives! ", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Blue("%s Takes a seat in the waiting room", client)
		default:
			//Buffered channel is full
			color.Red("Waiting room is full %s leaves", client)
		}

	} else {
		color.Red("The shop is already closed, so %s leaves", client)
	}

}
