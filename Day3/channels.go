package main

import (
	"fmt"
	"time"
)

func printCounts(label string, count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Println(label, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func main() {
	// Buffering available in channels
	// Synchronization available in channels: meaning one variable can be accessed by one goroutine at one time
	// Directional : means that channels can send only or receive only
	ch := make(chan int)
	go printCounts("goroutine", 5, ch)

	for value := range ch {
		fmt.Println("Main received: ", value)
	}
}
