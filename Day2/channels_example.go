package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func processOrder(orderId int, statusChannel chan Order) {
	time.Sleep(time.Second * 2)

	statusChannel <- Order{
		ID:     orderId,
		Status: "Completed",
	}
}

func main() {
	statusChannel := make(chan Order)

	for i := 1; i <= 5; i++ {
		go processOrder(i, statusChannel)
	}

	for i := 1; i <= 5; i++ {
		processedOrder := <-statusChannel
		fmt.Printf("Order %d processed: %d", processedOrder.ID, processedOrder.Status)
	}
}
