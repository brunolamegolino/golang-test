package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)

	}
}

func main() {
	channel := make(chan int)

	for i := 0; i < 100000; i++ {
		go worker(i, channel)
	}

	for i := 0; i < 1000000; i++ {
		channel <- i
	}

	time.Sleep(time.Second * 5)
}
