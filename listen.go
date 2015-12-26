package main

import (
	"fmt"
)

func work(myChan chan int, id int) {
	fmt.Println("Created channel number ", id)
	for {
		data := <-myChan
		fmt.Println("Processing data", data, " on channel id =", id)
	}
}

func receiver(instruction chan int, numOfWorkers int) {

	var myRoutes []chan int

	for i := 0; i < numOfWorkers; i++ {
		myRoutes = append(myRoutes, make(chan int))
		go work(myRoutes[i], i)
	}

	randomId := 0
	for {
		data := <-instruction
		fmt.Println("Sending to worker ", randomId)
		myRoutes[randomId] <- data
		if randomId == numOfWorkers-1 {
			randomId = 0
		}
		randomId = randomId + 1
	}
}

func test1(instruction chan int, wait chan bool, maxToSend int) {
	for i := 0; i < maxToSend; i++ {
		fmt.Println("Sending ", i)
		instruction <- i
	}
	// Telling wait that you could go home get some sleep.
}

func main() {

	fmt.Println("Starting the listener for incoming messages.")
	instruction := make(chan int)
	numOfChan := 1000
	maxToSend := 10000000

	wait := make(chan bool)
	go receiver(instruction, numOfChan)
	go test1(instruction, wait, maxToSend)
	<-wait
}
